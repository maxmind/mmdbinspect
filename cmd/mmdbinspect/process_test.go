package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	// We switch directories so that we don't have to deal with path separator
	// differences in the output.
	t.Chdir(testDataDir)

	tests := []struct {
		name      string
		useJSONL  bool
		networks  []string
		databases []string
		expected  string
	}{
		{
			name:      "single record YAML",
			useJSONL:  false,
			networks:  []string{"81.2.69.142"},
			databases: []string{"GeoIP2-Country-Test.mmdb"},
			expected:  "database_path:.*\nrequested_lookup: 81.2.69.142\nnetwork: 81.2.69.142/31\nrecord:\n",
		},
		{
			name:      "Single record JSONL",
			useJSONL:  true,
			networks:  []string{"81.2.69.142"},
			databases: []string{"GeoIP2-Country-Test.mmdb"},
			expected:  `{"database_path":"GeoIP2-Country-Test.mmdb","requested_lookup":"81.2.69.142","network":"81.2.69.142/31","record":{"continent":{"code":"EU","geoname_id":6255148,"names":{"de":"Europa","en":"Europe","es":"Europa","fr":"Europe","ja":"ヨーロッパ","pt-BR":"Europa","ru":"Европа","zh-CN":"欧洲"}},"country":{"geoname_id":2635167,"is_in_european_union":true,"iso_code":"GB","names":{"de":"Vereinigtes Königreich","en":"United Kingdom","es":"Reino Unido","fr":"Royaume-Uni","ja":"イギリス","pt-BR":"Reino Unido","ru":"Великобритания","zh-CN":"英国"}},"registered_country":{"geoname_id":6252001,"iso_code":"US","names":{"de":"USA","en":"United States","es":"Estados Unidos","fr":"États-Unis","ja":"アメリカ合衆国","pt-BR":"Estados Unidos","ru":"США","zh-CN":"美国"}}}}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := process(
				&buf,
				test.useJSONL,
				test.networks,
				test.databases,
				false,
				false,
				false,
			)

			require.NoError(t, err)

			if test.useJSONL {
				assert.JSONEq(t, test.expected, buf.String())
			} else {
				assert.Regexp(t, test.expected, buf.String())
			}
		})
	}
}
