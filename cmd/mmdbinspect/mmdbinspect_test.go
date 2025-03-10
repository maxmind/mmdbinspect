package main

import (
	"net/netip"
	"path/filepath"
	"testing"

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

func TestRecords(t *testing.T) {
	tests := []struct {
		name          string
		dbs           []string
		networks      []string
		expectRecords []record
		expectErr     string
	}{
		{
			name:     "multiple non-glob paths and multiple IPs",
			dbs:      []string{CityDBPath, CountryDBPath},
			networks: []string{"81.2.69.142", "8.8.8.8"},
			expectRecords: []record{
				{
					DatabasePath:    CityDBPath,
					RequestedLookup: "81.2.69.142",
					Network:         netip.MustParsePrefix("81.2.69.142/31"),
					Record:          city81_2_69_142,
				},
				{
					DatabasePath:    CountryDBPath,
					RequestedLookup: "81.2.69.142",
					Network:         netip.MustParsePrefix("81.2.69.142/31"),
					Record:          country81_2_69_142,
				},
			},
		},
		{
			name:     "glob path",
			dbs:      []string{filepath.Join(testDataDir, "GeoIP2-C*y-Test.mmdb")},
			networks: []string{"81.2.69.142"},
			expectRecords: []record{
				{
					DatabasePath:    CityDBPath,
					RequestedLookup: "81.2.69.142",
					Network:         netip.MustParsePrefix("81.2.69.142/31"),
					Record:          city81_2_69_142,
				},
				{
					DatabasePath:    CountryDBPath,
					RequestedLookup: "81.2.69.142",
					Network:         netip.MustParsePrefix("81.2.69.142/31"),
					Record:          country81_2_69_142,
				},
			},
		},
		{
			name:     "network missing from DB",
			dbs:      []string{CityDBPath},
			networks: []string{"10.0.0.0"},
		},
		{
			name:      "file does not exist",
			dbs:       []string{"does/not/exist.mmdb"},
			networks:  []string{"81.2.69.142"},
			expectErr: "does/not/exist.mmdb does not exist",
		},
		{
			name:      "invalid lookup IP",
			dbs:       []string{CityDBPath},
			networks:  []string{"81.2.69.342"},
			expectErr: "81.2.69.342 is not a valid IP address",
		},
		{
			name:      "invalid lookup network",
			dbs:       []string{CityDBPath},
			networks:  []string{"81.2.69.42/33"},
			expectErr: "81.2.69.42/33 is not a valid network",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var recs []record
			for record, err := range records(test.networks, test.dbs, false) {
				// For now, we don't test errors that happen half way through an
				// iteration. If we want to in the future, we will need to rework
				// this a bit.
				if test.expectErr == "" {
					require.NoError(t, err)
				} else {
					require.ErrorContains(t, err, test.expectErr)
				}

				if err == nil {
					recs = append(recs, *record)
				}
			}

			assert.Equal(t, test.expectRecords, recs)
		})
	}
}
