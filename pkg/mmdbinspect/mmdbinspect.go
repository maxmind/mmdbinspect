// Package mmdbinspect peeks at the contents of .mmdb files
package mmdbinspect

import (
	"errors"
	"fmt"
	"io/fs"
	"iter"
	"net/netip"
	"os"
	"path/filepath"
	"strings"

	"github.com/oschwald/maxminddb-golang/v2"
)

// Record holds the records for a lookup in a database.
type Record struct {
	DatabasePath    string       `json:"database_path"`
	RequestedLookup string       `json:"requested_lookup"`
	Network         netip.Prefix `json:"network"`
	Record          any          `json:"record"`
}

// Records returns an iterator over the records for the networks and
// databases provided.
func Records(
	networks, databases []string,
	includeAliasedNetworks bool,
) iter.Seq2[*Record, error] {
	return func(yield func(*Record, error) bool) {
		for _, glob := range databases {
			matches, err := filepath.Glob(glob)
			if err != nil {
				yield(nil, fmt.Errorf("invalid file path or glob %q: %w", glob, err))
				return
			}
			for _, path := range matches {
				reader, err := openDB(path)
				if err != nil {
					yield(nil, fmt.Errorf("could not open database %q: %w", path, err))
					return
				}

				for _, thisNetwork := range networks {
					baseRecord := Record{
						DatabasePath:    path,
						RequestedLookup: thisNetwork,
					}
					ok := recordsForNetwork(reader, includeAliasedNetworks, baseRecord, yield)
					if !ok {
						return
					}
				}
				_ = reader.Close()
			}
		}
	}
}

// recordsForNetwork fetches mmdb records inside a given network. If an IP
// address is provided without a prefix length, it will be treated as a
// network containing a single address (i.e., /32 for IPv4 and /128 for IPv6).
func recordsForNetwork(
	reader *maxminddb.Reader,
	includeAliasedNetworks bool,
	record Record,
	yield func(*Record, error) bool,
) bool {
	lookupNetwork := record.RequestedLookup

	if !strings.Contains(lookupNetwork, "/") {
		if strings.Count(lookupNetwork, ":") >= 2 {
			lookupNetwork = lookupNetwork + "/128"
		} else {
			lookupNetwork = lookupNetwork + "/32"
		}
	}

	network, err := netip.ParsePrefix(lookupNetwork)
	if err != nil {
		yield(nil, fmt.Errorf("%v is not a valid network or IP address", record.RequestedLookup))
		return false
	}

	var opts []maxminddb.NetworksOption
	if includeAliasedNetworks {
		opts = append(opts, maxminddb.IncludeAliasedNetworks)
	}

	for res := range reader.NetworksWithin(network, opts...) {
		record.Network = res.Prefix()

		err := res.Decode(&record.Record)
		if err != nil {
			yield(nil, fmt.Errorf("could not get next network: %w", err))
			return false
		}

		ok := yield(&record, nil)
		if !ok {
			return false
		}
	}

	return true
}

// openDB returns a maxminddb.Reader.
func openDB(path string) (*maxminddb.Reader, error) {
	_, err := os.Stat(path)

	if errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("%v does not exist", path)
	}
	if err != nil {
		return nil, fmt.Errorf("stating %s: %w", path, err)
	}

	db, err := maxminddb.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%v could not be opened: %w", path, err)
	}

	return db, nil
}
