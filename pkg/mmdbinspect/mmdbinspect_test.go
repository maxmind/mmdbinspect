package mmdbinspect

import (
	"testing"

	"github.com/oschwald/maxminddb-golang"
	"github.com/stretchr/testify/assert"
)

const CityDBPath = "../../test/data/test-data/GeoIP2-City-Test.mmdb"
const CountryDBPath = "../../test/data/test-data/GeoIP2-Country-Test.mmdb"

func TestOpenDB(t *testing.T) {
	assert := assert.New(t)

	assert.FileExists(CityDBPath, "database exists")

	var reader, err = OpenDB(CityDBPath)
	assert.NoError(err, "no open error")
	assert.IsType(maxminddb.Reader{}, *reader)

	reader, err = OpenDB("foo/bar/baz")
	assert.Error(err, "open error when file does not exist")
	assert.Nil(reader)
	assert.Equal(
		"foo/bar/baz does not exist",
		err.Error(),
	)

	reader, err = OpenDB("../../test/data/test-data/README.md")
	assert.Error(err)
	assert.Contains(err.Error(), "README.md could not be opened: error opening database: invalid MaxMind DB file")
	assert.Nil(reader)

	reader, err = OpenDB("../../test/data/test-data/GeoIP2-City-Test-Invalid-Node-Count.mmdb")
	assert.Error(err)
	assert.Contains(err.Error(), "invalid metadata")
	assert.Nil(reader)

	if reader != nil {
		reader.Close()
	}
}

func TestRecordsForNetwork(t *testing.T) {
	assert := assert.New(t)
	var reader, _ = OpenDB(CityDBPath) // ipv6 database

	var records, err = RecordsForNetwork(*reader, "81.2.69.142")
	assert.NoError(err, "no error on lookup of 81.2.69.142")
	assert.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, "81.2.69.0/24")
	assert.NoError(err, "no error on lookup of 81.2.69.0/24")
	assert.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, "10.255.255.255/29")
	assert.NoError(err, "got no error when IP not found")
	assert.Nil(records, "no records returned for 10.255.255.255/29")

	records, err = RecordsForNetwork(*reader, "X.X.Y.Z")
	assert.Error(err, "got an error")
	assert.Nil(records, "no records returned for X.X.Y.Z")
	assert.Equal("X.X.Y.Z is not a valid IP address", err.Error())

	reader.Close()
}

func TestRecordToString(t *testing.T) {
	assert := assert.New(t)

	var reader, _ = OpenDB(CityDBPath)
	var records, _ = RecordsForNetwork(*reader, "81.2.69.142")
	var prettyJSON, err = RecordToString(records)

	assert.NoError(err, "no error on stringification")
	assert.NotNil(prettyJSON, "records stringified")
	assert.Contains(prettyJSON, "London")
	assert.Contains(prettyJSON, "2643743")

	reader.Close()
}

func TestAggregatedRecords(t *testing.T) {
	assert := assert.New(t)

	dbs := []string{CityDBPath, CountryDBPath}
	networks := []string{"81.2.69.142", "8.8.8.8"}
	results, err := AggregatedRecords(networks, dbs)

	assert.NoError(err)
	assert.NotNil(results)
}
