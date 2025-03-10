// Package mmdbinspect peeks at the contents of .mmdb files
package mmdbinspect

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/netip"
	"os"
	"path/filepath"
	"strings"

	"github.com/oschwald/maxminddb-golang/v2"
)

// RecordForNetwork holds a network and the corresponding record.
type RecordForNetwork struct {
	Network netip.Prefix
	Record  any
}

// RecordSet holds the records for a lookup in a database.
type RecordSet struct {
	Database string
	Records  []RecordForNetwork
	Lookup   string
}

// OpenDB returns a maxminddb.Reader.
func OpenDB(path string) (*maxminddb.Reader, error) {
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

// RecordsForNetwork fetches mmdb records inside a given network. If an IP
// address is provided without a prefix length, it will be treated as a
// network containing a single address (i.e., /32 for IPv4 and /128 for IPv6).
func RecordsForNetwork(
	reader maxminddb.Reader,
	includeAliasedNetworks bool,
	maybeNetwork string,
) ([]RecordForNetwork, error) {
	lookupNetwork := maybeNetwork

	if !strings.Contains(lookupNetwork, "/") {
		if strings.Count(maybeNetwork, ":") >= 2 {
			lookupNetwork = maybeNetwork + "/128"
		} else {
			lookupNetwork = maybeNetwork + "/32"
		}
	}

	network, err := netip.ParsePrefix(lookupNetwork)
	if err != nil {
		return nil, fmt.Errorf("%v is not a valid IP address", maybeNetwork)
	}

	var opts []maxminddb.NetworksOption
	if includeAliasedNetworks {
		opts = append(opts, maxminddb.IncludeAliasedNetworks)
	}

	var found []RecordForNetwork

	for res := range reader.NetworksWithin(network, opts...) {
		var record any

		err := res.Decode(&record)
		if err != nil {
			return nil, fmt.Errorf("could not get next network: %w", err)
		}

		found = append(found, RecordForNetwork{res.Prefix(), record})
	}

	return found, nil
}

// AggregatedRecords returns the aggregated records for the networks and
// databases provided.
func AggregatedRecords(
	networks, databases []string,
	includeAliasedNetworks bool,
) ([]RecordSet, error) {
	var recordSets []RecordSet

	for _, glob := range databases {
		matches, err := filepath.Glob(glob)
		if err != nil {
			return nil, fmt.Errorf("invalid file path or glob %q: %w", glob, err)
		}
		for _, path := range matches {
			reader, err := OpenDB(path)
			if err != nil {
				return nil, fmt.Errorf("could not open database %q: %w", path, err)
			}

			for _, thisNetwork := range networks {
				records, err := RecordsForNetwork(*reader, includeAliasedNetworks, thisNetwork)
				if err != nil {
					_ = reader.Close()
					return nil, fmt.Errorf("could not get records from db %q: %w", path, err)
				}

				set := RecordSet{path, records, thisNetwork}
				recordSets = append(recordSets, set)
			}
			_ = reader.Close()
		}
	}

	return recordSets, nil
}

// RecordToString converts an mmdb record into a JSON-formatted string.
func RecordToString(record []RecordSet) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false) // don't escape ampersands and angle brackets
	enc.SetIndent("", "    ")

	err := enc.Encode(record)
	if err != nil {
		return "", errors.New("could not convert record to string")
	}

	return buf.String(), nil
}
