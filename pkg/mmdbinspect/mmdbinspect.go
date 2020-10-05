// Package mmdbinspect peeks at the contents of .mmdb files
package mmdbinspect

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/oschwald/maxminddb-golang"
	"github.com/pkg/errors"
)

// RecordForNetwork holds a network and the corresponding record.
type RecordForNetwork struct {
	Network string
	Record  interface{}
}

// RecordSet holds the records for a lookup in a database.
type RecordSet struct {
	Database string
	Records  interface{}
	Lookup   string
}

// OpenDB returns a maxminddb.Reader
func OpenDB(path string) (*maxminddb.Reader, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return nil, errors.Errorf("%v does not exist", path)
	} else if err != nil {
		return nil, err
	}

	db, err := maxminddb.Open(path)
	if err != nil {
		return nil, errors.WithMessagef(err, "%v could not be opened", path)
	}

	return db, nil
}

// RecordsForNetwork fetches mmdb records inside a given network.  If an
// address is provided without a netmask a /32 will be inferred for v4
// addresses and a /128 will be inferred for v6 addresses.
func RecordsForNetwork(reader maxminddb.Reader, maybeNetwork string) (interface{}, error) {
	lookupNetwork := maybeNetwork

	if !strings.Contains(lookupNetwork, "/") {
		if strings.Count(maybeNetwork, ":") >= 2 {
			lookupNetwork = maybeNetwork + "/128"
		} else {
			lookupNetwork = maybeNetwork + "/32"
		}
	}

	_, network, err := net.ParseCIDR(lookupNetwork)
	if err != nil {
		return nil, errors.Errorf("%v is not a valid IP address", maybeNetwork)
	}

	n := reader.NetworksWithin(network)

	var found []interface{}

	for n.Next() {
		var record interface{}
		address, err := n.Network(&record)
		if err != nil {
			return nil, errors.WithMessagef(err, "Could not get next network")
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
func AggregatedRecords(networks, databases []string) (interface{}, error) {
	recordSets := make([]RecordSet, 0)

	for _, path := range databases {
		reader, err := OpenDB(path)
		if err != nil {
			return nil, errors.WithMessagef(err, "could not open database %v", path)
		}
		defer reader.Close()

		for _, thisNetwork := range networks {
			var records interface{}
			records, err = RecordsForNetwork(*reader, thisNetwork)

			if err != nil {
				return nil, errors.WithMessagef(err, "could not get records from db %v", path)
			}

			set := RecordSet{path, records, thisNetwork}
			recordSets = append(recordSets, set)
		}
	}

	return recordSets, nil
}

// RecordToString converts an mmdb record into a JSON-formatted string
func RecordToString(record interface{}) (string, error) {
	json, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		return "", errors.Errorf("Could not convert record to string")
	}

	return fmt.Sprintf("%v", string(json)), nil
}
