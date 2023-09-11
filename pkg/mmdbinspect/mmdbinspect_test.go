package mmdbinspect

import (
	"testing"

	"github.com/oschwald/maxminddb-golang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	CityDBPath    = "../../test/data/test-data/GeoIP2-City-Test.mmdb"
	CountryDBPath = "../../test/data/test-data/GeoIP2-Country-Test.mmdb"
)

func TestOpenDB(t *testing.T) {
	a := assert.New(t)

	a.FileExists(CityDBPath, "database exists")

	reader, err := OpenDB(CityDBPath)
	a.NoError(err, "no open error")
	a.IsType(maxminddb.Reader{}, *reader)

	reader, err = OpenDB("foo/bar/baz")
	a.Error(err, "open error when file does not exist")
	a.Nil(reader)
	a.Equal(
		"foo/bar/baz does not exist",
		err.Error(),
	)

	reader, err = OpenDB("../../test/data/test-data/README.md")
	a.Error(err)
	a.Contains(err.Error(), "README.md could not be opened: error opening database: invalid MaxMind DB file")
	a.Nil(reader)

	reader, err = OpenDB("../../test/data/test-data/GeoIP2-City-Test-Invalid-Node-Count.mmdb")
	a.Error(err)
	a.Contains(err.Error(), "invalid metadata")
	a.Nil(reader)

	if reader != nil {
		require.NoError(t, reader.Close())
	}
}

func TestRecordsForNetwork(t *testing.T) {
	a := assert.New(t)
	reader, err := OpenDB(CityDBPath) // ipv6 database
	a.NoError(err, "no open error")

	records, err := RecordsForNetwork(*reader, false, "81.2.69.142")
	a.NoError(err, "no error on lookup of 81.2.69.142")
	a.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, false, "81.2.69.0/24")
	a.NoError(err, "no error on lookup of 81.2.69.0/24")
	a.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, false, "10.255.255.255/29")
	a.NoError(err, "got no error when IP not found")
	a.Nil(records, "no records returned for 10.255.255.255/29")

	records, err = RecordsForNetwork(*reader, false, "X.X.Y.Z")
	a.Error(err, "got an error")
	a.Nil(records, "no records returned for X.X.Y.Z")
	a.Equal("X.X.Y.Z is not a valid IP address", err.Error())

	require.NoError(t, reader.Close())
}

func TestRecordToString(t *testing.T) {
	a := assert.New(t)

	reader, err := OpenDB(CityDBPath)
	a.NoError(err, "no open error")
	records, err := RecordsForNetwork(*reader, false, "81.2.69.142")
	a.NoError(err, "no RecordsForNetwork error")
	prettyJSON, err := RecordToString(records)

	a.NoError(err, "no error on stringification")
	a.NotNil(prettyJSON, "records stringified")
	a.Contains(prettyJSON, "London")
	a.Contains(prettyJSON, "2643743")

	require.NoError(t, reader.Close())
}

func TestAggregatedRecords(t *testing.T) {
	a := assert.New(t)

	dbs := []string{CityDBPath, CountryDBPath}
	networks := []string{"81.2.69.142", "8.8.8.8"}
	results, err := AggregatedRecords(networks, dbs, false)

	a.NoError(err)
	a.NotNil(results)
}
