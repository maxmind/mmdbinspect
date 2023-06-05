// Package mmdbinspect peeks at the contents of .mmdb files
package mmdbinspect

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/oschwald/maxminddb-golang"
)

// RecordForNetwork holds a network and the corresponding record.
type RecordForNetwork struct {
	Network string
	Record  any
}

// RecordSet holds the records for a lookup in a database.
type RecordSet struct {
	Database string
	Records  any
	Lookup   string
}

// OpenDB returns a maxminddb.Reader.
func OpenDB(path string) (*maxminddb.Reader, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%v does not exist", path)
	} else if err != nil {
		return nil, err
	}

	db, err := maxminddb.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%v could not be opened: %w", path, err)
	}

	return db, nil
}

// RecordsForNetwork fetches mmdb records inside a given network.  If an
// address is provided without a netmask a /32 will be inferred for v4
// addresses and a /128 will be inferred for v6 addresses.
func RecordsForNetwork(reader maxminddb.Reader, maybeNetwork string) (any, error) {
	lookupNetwork := maybeNetwork

	if !strings.Contains(lookupNetwork, "/") {
		if strings.Count(maybeNetwork, ":") >= 2 {
			lookupNetwork = maybeNetwork + "/128"
		} else {
			lookupNetwork = maybeNetwork + "/32"
		}
	}

	//nolint:forbidigo // preexisting
	_, network, err := net.ParseCIDR(lookupNetwork)
	if err != nil {
		return nil, fmt.Errorf("%v is not a valid IP address", maybeNetwork)
	}

	n := reader.NetworksWithin(network)

	var found []any

	for n.Next() {
		var record any
		address, err := n.Network(&record)
		if err != nil {
			return nil, fmt.Errorf("could not get next network: %w", err)
		}

		found = append(found, RecordForNetwork{address.String(), record})
	}

	if n.Err() != nil {
		return nil, n.Err()
	}

	return found, nil
}

// AggregatedRecords returns the aggregated records for the networks and
// databases provided.
func AggregatedRecords(networks, databases []string) (any, error) {
	recordSets := make([]RecordSet, 0)

	for _, path := range databases {
		reader, err := OpenDB(path)
		if err != nil {
			return nil, fmt.Errorf("could not open database %v: %w", path, err)
		}

		for _, thisNetwork := range networks {
			var records any
			records, err = RecordsForNetwork(*reader, thisNetwork)

			if err != nil {
				_ = reader.Close()
				return nil, fmt.Errorf("could not get records from db %v: %w", path, err)
			}

			set := RecordSet{path, records, thisNetwork}
			recordSets = append(recordSets, set)
		}
		_ = reader.Close()
	}

	return recordSets, nil
}

// RecordToString converts an mmdb record into a JSON-formatted string.
func RecordToString(record any) (string, error) {
	j, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		return "", errors.New("could not convert record to string")
	}

	return string(j), nil
}
