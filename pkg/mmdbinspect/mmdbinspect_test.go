package mmdbinspect

import (
	"net/netip"
	"path/filepath"
	"testing"

	"github.com/oschwald/maxminddb-golang/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDataDir = "../../test/data/test-data/"
)

var (
	CityDBPath    = filepath.Join(testDataDir, "GeoIP2-City-Test.mmdb")
	CountryDBPath = filepath.Join(testDataDir, "GeoIP2-Country-Test.mmdb")
	ISPDBPath     = filepath.Join(testDataDir, "GeoIP2-ISP-Test.mmdb")
)

func TestOpenDB(t *testing.T) {
	a := assert.New(t)

	a.FileExists(CityDBPath, "database exists")

	reader, err := OpenDB(CityDBPath)
	require.NoError(t, err, "no open error")
	a.IsType(maxminddb.Reader{}, *reader)

	reader, err = OpenDB("foo/bar/baz")
	require.Error(t, err, "open error when file does not exist")
	a.Nil(reader)
	a.Equal(
		"foo/bar/baz does not exist",
		err.Error(),
	)

	reader, err = OpenDB("../../test/data/test-data/README.md")
	require.Error(t, err)
	a.Contains(err.Error(), "README.md could not be opened: error opening database: invalid MaxMind DB file")
	a.Nil(reader)

	reader, err = OpenDB("../../test/data/test-data/GeoIP2-City-Test-Invalid-Node-Count.mmdb")
	require.Error(t, err)
	a.Contains(err.Error(), "invalid metadata")
	a.Nil(reader)

	if reader != nil {
		require.NoError(t, reader.Close())
	}
}

func TestRecordsForNetwork(t *testing.T) {
	a := assert.New(t)
	reader, err := OpenDB(CityDBPath) // ipv6 database
	require.NoError(t, err, "no open error")

	records, err := RecordsForNetwork(*reader, false, "81.2.69.142")
	require.NoError(t, err, "no error on lookup of 81.2.69.142")
	a.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, false, "81.2.69.0/24")
	require.NoError(t, err, "no error on lookup of 81.2.69.0/24")
	a.NotNil(records, "records returned")

	records, err = RecordsForNetwork(*reader, false, "10.255.255.255/29")
	require.NoError(t, err, "got no error when IP not found")
	a.Nil(records, "no records returned for 10.255.255.255/29")

	records, err = RecordsForNetwork(*reader, false, "X.X.Y.Z")
	require.Error(t, err, "got an error")
	a.Nil(records, "no records returned for X.X.Y.Z")
	a.Equal("X.X.Y.Z is not a valid IP address", err.Error())

	require.NoError(t, reader.Close())
}

func TestRecordToString(t *testing.T) {
	a := assert.New(t)

	reader, err := OpenDB(CityDBPath)
	require.NoError(t, err, "no open error")
	records, err := RecordsForNetwork(*reader, false, "81.2.69.142")
	require.NoError(t, err, "no RecordsForNetwork error")
	prettyJSON, err := RecordToString(records)

	require.NoError(t, err, "no error on stringification")
	a.NotNil(prettyJSON, "records stringified")
	a.Contains(prettyJSON, "London")
	a.Contains(prettyJSON, "2643743")

	require.NoError(t, reader.Close())
}

// TestRecordToStringEscaping tests that certain HTML-related characters are not
// escaped in the JSON output.
func TestRecordToStringEscaping(t *testing.T) {
	a := assert.New(t)

	reader, err := OpenDB(ISPDBPath)
	require.NoError(t, err, "no open error")
	records, err := RecordsForNetwork(*reader, false, "206.16.137.0/24")
	require.NoError(t, err, "no RecordsForNetwork error")
	prettyJSON, err := RecordToString(records)

	require.NoError(t, err, "no error on stringification")
	a.NotNil(prettyJSON, "records stringified")
	a.Contains(prettyJSON, "AT&T Synaptic Cloud Hosting")

	require.NoError(t, reader.Close())
}

var city81_2_69_142 = map[string]any{
	"city": map[string]any{
		"geoname_id": uint64(2643743),
		"names": map[string]any{
			"de":    "London",
			"en":    "London",
			"es":    "Londres",
			"fr":    "Londres",
			"ja":    "ロンドン",
			"pt-BR": "Londres",
			"ru":    "Лондон",
		},
	},
	"continent": map[string]any{
		"code":       "EU",
		"geoname_id": uint64(6255148),
		"names": map[string]any{
			"de":    "Europa",
			"en":    "Europe",
			"es":    "Europa",
			"fr":    "Europe",
			"ja":    "ヨーロッパ",
			"pt-BR": "Europa",
			"ru":    "Европа",
			"zh-CN": "欧洲",
		},
	},
	"country": map[string]any{
		"geoname_id":           uint64(2635167),
		"is_in_european_union": true,
		"iso_code":             "GB",
		"names": map[string]any{
			"de":    "Vereinigtes Königreich",
			"en":    "United Kingdom",
			"es":    "Reino Unido",
			"fr":    "Royaume-Uni",
			"ja":    "イギリス",
			"pt-BR": "Reino Unido",
			"ru":    "Великобритания",
			"zh-CN": "英国",
		},
	},
	"location": map[string]any{
		"accuracy_radius": uint64(10),
		"latitude":        51.5142,
		"longitude":       -0.0931,
		"time_zone":       "Europe/London",
	},
	"registered_country": map[string]any{
		"geoname_id": uint64(6252001),
		"iso_code":   "US",
		"names": map[string]any{
			"de":    "USA",
			"en":    "United States",
			"es":    "Estados Unidos",
			"fr":    "États-Unis",
			"ja":    "アメリカ合衆国",
			"pt-BR": "Estados Unidos",
			"ru":    "США",
			"zh-CN": "美国",
		},
	},
	"subdivisions": []any{map[string]any{
		"geoname_id": uint64(6269131),
		"iso_code":   "ENG",
		"names": map[string]any{
			"en":    "England",
			"es":    "Inglaterra",
			"fr":    "Angleterre",
			"pt-BR": "Inglaterra",
		},
	}},
}

var country81_2_69_142 = map[string]any{
	"continent": map[string]any{
		"code":       "EU",
		"geoname_id": uint64(6255148),
		"names": map[string]any{
			"de":    "Europa",
			"en":    "Europe",
			"es":    "Europa",
			"fr":    "Europe",
			"ja":    "ヨーロッパ",
			"pt-BR": "Europa",
			"ru":    "Европа",
			"zh-CN": "欧洲",
		},
	},
	"country": map[string]any{
		"geoname_id":           uint64(2635167),
		"is_in_european_union": true,
		"iso_code":             "GB",
		"names": map[string]any{
			"de":    "Vereinigtes Königreich",
			"en":    "United Kingdom",
			"es":    "Reino Unido",
			"fr":    "Royaume-Uni",
			"ja":    "イギリス",
			"pt-BR": "Reino Unido",
			"ru":    "Великобритания",
			"zh-CN": "英国",
		},
	},
	"registered_country": map[string]any{
		"geoname_id": uint64(6252001),
		"iso_code":   "US",
		"names": map[string]any{
			"de":    "USA",
			"en":    "United States",
			"es":    "Estados Unidos",
			"fr":    "États-Unis",
			"ja":    "アメリカ合衆国",
			"pt-BR": "Estados Unidos",
			"ru":    "США",
			"zh-CN": "美国",
		},
	},
}

func TestAggregatedRecords(t *testing.T) {
	tests := []struct {
		name     string
		dbs      []string
		networks []string
		expected []RecordSet
	}{
		{
			name:     "multiple non-glob paths and multiple IPs",
			dbs:      []string{CityDBPath, CountryDBPath},
			networks: []string{"81.2.69.142", "8.8.8.8"},
			expected: []RecordSet{
				{
					Database: CityDBPath,
					Records: []any{
						RecordForNetwork{
							Network: netip.MustParsePrefix("81.2.69.142/31"),
							Record:  city81_2_69_142,
						},
					},
					Lookup: "81.2.69.142",
				},
				{
					Database: CityDBPath,
					Records:  []any(nil),
					Lookup:   "8.8.8.8",
				},
				{
					Database: CountryDBPath,
					Records: []any{
						RecordForNetwork{
							Network: netip.MustParsePrefix("81.2.69.142/31"),
							Record:  country81_2_69_142,
						},
					},
					Lookup: "81.2.69.142",
				},
				{
					Database: CountryDBPath,
					Records:  []any(nil),
					Lookup:   "8.8.8.8",
				},
			},
		},
		{
			name:     "glob path",
			dbs:      []string{filepath.Join(testDataDir, "GeoIP2-C*y-Test.mmdb")},
			networks: []string{"81.2.69.142"},
			expected: []RecordSet{
				{
					Database: CityDBPath,
					Records: []any{
						RecordForNetwork{
							Network: netip.MustParsePrefix("81.2.69.142/31"),
							Record:  city81_2_69_142,
						},
					},
					Lookup: "81.2.69.142",
				},
				{
					Database: CountryDBPath,
					Records: []any{
						RecordForNetwork{
							Network: netip.MustParsePrefix("81.2.69.142/31"),
							Record:  country81_2_69_142,
						},
					},
					Lookup: "81.2.69.142",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results, err := AggregatedRecords(test.networks, test.dbs, false)
			require.NoError(t, err)

			assert.Equal(t, test.expected, results)
		})
	}
}
