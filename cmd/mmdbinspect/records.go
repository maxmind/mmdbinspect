package main

import (
	"errors"
	"fmt"
	"io/fs"
	"iter"
	"net/netip"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/oschwald/maxminddb-golang/v2"
)

// record holds the records for a lookup in a database. Note that
// the order here affects the output order.
type record struct {
	DatabasePath    string       `json:"database_path"`
	BuildTime       *time.Time   `json:"build_time,omitempty"`
	RequestedLookup string       `json:"requested_lookup"`
	Network         netip.Prefix `json:"network"`
	Record          any          `json:"record,omitempty"`
}

// records returns an iterator over the records for the networks and
// databases provided.
func records(
	networks, databases []string,
	includeAliasedNetworks,
	includeBuildTime,
	includeNetworksWithoutData,
	includeEmptyValues bool,
) iter.Seq2[*record, error] {
	var opts []maxminddb.NetworksOption
	if includeAliasedNetworks {
		opts = append(opts, maxminddb.IncludeAliasedNetworks())
	}
	if includeNetworksWithoutData {
		opts = append(opts, maxminddb.IncludeNetworksWithoutData())
	}
	if !includeEmptyValues {
		opts = append(opts, maxminddb.SkipEmptyValues())
	}

	return func(yield func(*record, error) bool) {
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

				var buildTime *time.Time
				if includeBuildTime {
					//nolint:gosec // BuildEpoch is a Unix timestamp, won't overflow
					t := time.Unix(int64(reader.Metadata.BuildEpoch), 0).UTC()
					buildTime = &t
				}

				for _, thisNetwork := range networks {
					baseRecord := record{
						DatabasePath:    path,
						BuildTime:       buildTime,
						RequestedLookup: thisNetwork,
					}
					ok := recordsForNetwork(reader, opts, baseRecord, yield)
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
	opts []maxminddb.NetworksOption,
	record record,
	yield func(*record, error) bool,
) bool {
	var err error
	var network netip.Prefix
	if strings.Contains(record.RequestedLookup, "/") {
		network, err = netip.ParsePrefix(record.RequestedLookup)
		if err != nil {
			yield(nil, fmt.Errorf("%s is not a valid network", record.RequestedLookup))
			return false
		}
	} else {
		addr, err := netip.ParseAddr(record.RequestedLookup)
		if err != nil {
			yield(nil, fmt.Errorf("%s is not a valid IP address", record.RequestedLookup))
			return false
		}

		bits := 32
		if addr.Is6() {
			bits = 128
		}

		network = netip.PrefixFrom(addr, bits)
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
