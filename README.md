Inspect `.mmdb` databases

```bash
cd cmd/mmdbinspect
go build
./mmdbinspect --db ../../GeoLite2-Country.mmdb --db ../../dbs/GeoLite2-City.mmdb 130.113.64.30/24 0:0:0:0:0:ffff:8064:a678
```

Yields:

```bash
[
    {
        "Database": "../../GeoLite2-Country.mmdb",
        "Records": [
            {
                "Network": "130.113.64.0/16",
                "Record": {
                    "continent": {
                        "code": "NA",
                        "geoname_id": 6255149,
                        "names": {
                            "de": "Nordamerika",
                            "en": "North America",
                            "es": "Norteamérica",
                            "fr": "Amérique du Nord",
                            "ja": "北アメリカ",
                            "pt-BR": "América do Norte",
                            "ru": "Северная Америка",
                            "zh-CN": "北美洲"
                        }
                    },
                    "country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "registered_country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    }
                }
            }
        ],
        "Lookup": "130.113.64.30/24"
    },
    {
        "Database": "../../GeoLite2-Country.mmdb",
        "Records": [
            {
                "Network": "128.100.166.120/16",
                "Record": {
                    "continent": {
                        "code": "NA",
                        "geoname_id": 6255149,
                        "names": {
                            "de": "Nordamerika",
                            "en": "North America",
                            "es": "Norteamérica",
                            "fr": "Amérique du Nord",
                            "ja": "北アメリカ",
                            "pt-BR": "América do Norte",
                            "ru": "Северная Америка",
                            "zh-CN": "北美洲"
                        }
                    },
                    "country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "registered_country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    }
                }
            }
        ],
        "Lookup": "0:0:0:0:0:ffff:8064:a678"
    },
    {
        "Database": "../../dbs/GeoLite2-City.mmdb",
        "Records": [
            {
                "Network": "130.113.64.0/16",
                "Record": {
                    "city": {
                        "geoname_id": 5969782,
                        "names": {
                            "en": "Hamilton",
                            "ja": "ハミルトン",
                            "ru": "Гамильтон",
                            "zh-CN": "哈密尔顿"
                        }
                    },
                    "continent": {
                        "code": "NA",
                        "geoname_id": 6255149,
                        "names": {
                            "de": "Nordamerika",
                            "en": "North America",
                            "es": "Norteamérica",
                            "fr": "Amérique du Nord",
                            "ja": "北アメリカ",
                            "pt-BR": "América do Norte",
                            "ru": "Северная Америка",
                            "zh-CN": "北美洲"
                        }
                    },
                    "country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "location": {
                        "accuracy_radius": 5,
                        "latitude": 43.2642,
                        "longitude": -79.9143,
                        "time_zone": "America/Toronto"
                    },
                    "postal": {
                        "code": "L8S"
                    },
                    "registered_country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "subdivisions": [
                        {
                            "geoname_id": 6093943,
                            "iso_code": "ON",
                            "names": {
                                "en": "Ontario",
                                "fr": "Ontario",
                                "ja": "オンタリオ州",
                                "pt-BR": "Ontário",
                                "ru": "Онтарио",
                                "zh-CN": "安大略"
                            }
                        }
                    ]
                }
            }
        ],
        "Lookup": "130.113.64.30/24"
    },
    {
        "Database": "../../dbs/GeoLite2-City.mmdb",
        "Records": [
            {
                "Network": "128.100.166.120/16",
                "Record": {
                    "city": {
                        "geoname_id": 6167865,
                        "names": {
                            "de": "Toronto",
                            "en": "Toronto",
                            "es": "Toronto",
                            "fr": "Toronto",
                            "ja": "トロント",
                            "pt-BR": "Toronto",
                            "ru": "Торонто",
                            "zh-CN": "多伦多"
                        }
                    },
                    "continent": {
                        "code": "NA",
                        "geoname_id": 6255149,
                        "names": {
                            "de": "Nordamerika",
                            "en": "North America",
                            "es": "Norteamérica",
                            "fr": "Amérique du Nord",
                            "ja": "北アメリカ",
                            "pt-BR": "América do Norte",
                            "ru": "Северная Америка",
                            "zh-CN": "北美洲"
                        }
                    },
                    "country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "location": {
                        "accuracy_radius": 5,
                        "latitude": 43.6638,
                        "longitude": -79.3999,
                        "time_zone": "America/Toronto"
                    },
                    "postal": {
                        "code": "M5S"
                    },
                    "registered_country": {
                        "geoname_id": 6251999,
                        "iso_code": "CA",
                        "names": {
                            "de": "Kanada",
                            "en": "Canada",
                            "es": "Canadá",
                            "fr": "Canada",
                            "ja": "カナダ",
                            "pt-BR": "Canadá",
                            "ru": "Канада",
                            "zh-CN": "加拿大"
                        }
                    },
                    "subdivisions": [
                        {
                            "geoname_id": 6093943,
                            "iso_code": "ON",
                            "names": {
                                "en": "Ontario",
                                "fr": "Ontario",
                                "ja": "オンタリオ州",
                                "pt-BR": "Ontário",
                                "ru": "Онтарио",
                                "zh-CN": "安大略"
                            }
                        }
                    ]
                }
            }
        ],
        "Lookup": "0:0:0:0:0:ffff:8064:a678"
    }
]
```

Or, pipe it to `jq`:

```bash
./mmdbinspect \
--db ../../GeoLite2-Country.mmdb \
--db ../../dbs/GeoLite2-City.mmdb \
130.113.64.30/24 0:0:0:0:0:ffff:8064:a678 \
| jq '.[] | {database: .Database, Lookup: .Lookup, Network: .Records[].Network, Country: .Records[].Record.country.names.en,  City: .Records[].Record.city.names.en,}' 
```

Yields:
```bash
Country: .Records[].Record.country.names.en,  City: .Records[].Record.city.names.en,}'
{
  "database": "../../GeoLite2-Country.mmdb",
  "Lookup": "130.113.64.30/24",
  "Network": "130.113.64.0/16",
  "Country": "Canada",
  "City": null
}
{
  "database": "../../GeoLite2-Country.mmdb",
  "Lookup": "0:0:0:0:0:ffff:8064:a678",
  "Network": "128.100.166.120/16",
  "Country": "Canada",
  "City": null
}
{
  "database": "../../dbs/GeoLite2-City.mmdb",
  "Lookup": "130.113.64.30/24",
  "Network": "130.113.64.0/16",
  "Country": "Canada",
  "City": "Hamilton"
}
{
  "database": "../../dbs/GeoLite2-City.mmdb",
  "Lookup": "0:0:0:0:0:ffff:8064:a678",
  "Network": "128.100.166.120/16",
  "Country": "Canada",
  "City": "Toronto"
}
```
