Inspect `.mmdb` databases

```bash
go run cmd/mmdbinspect/main.go --db dbs/GeoLite2-City.mmdb --network 69.158.16.0/32
```

```bash
go run cmd/mmdbinspect/main.go \
--db dbs/GeoLite2-Country.mmdb  \
--db dbs/GeoLite2-City.mmdb \
--network 69.158.16.0/32    \
--network 8.0.17.0/24
```

Yields:

```bash
[
    {
        "Database": "dbs/GeoLite2-Country.mmdb",
        "Records": null
    },
    {
        "Database": "dbs/GeoLite2-City.mmdb",
        "Records": [
            {
                "Network": "69.158.16.0/21",
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
                        "accuracy_radius": 10,
                        "latitude": 43.6661,
                        "longitude": -79.529,
                        "time_zone": "America/Toronto"
                    },
                    "postal": {
                        "code": "M9A"
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
        ]
    },
    {
        "Database": "dbs/GeoLite2-Country.mmdb",
        "Records": [
            {
                "Network": "8.0.17.0/24",
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
                        "geoname_id": 6252001,
                        "iso_code": "US",
                        "names": {
                            "de": "USA",
                            "en": "United States",
                            "es": "Estados Unidos",
                            "fr": "États-Unis",
                            "ja": "アメリカ合衆国",
                            "pt-BR": "Estados Unidos",
                            "ru": "США",
                            "zh-CN": "美国"
                        }
                    },
                    "registered_country": {
                        "geoname_id": 6252001,
                        "iso_code": "US",
                        "names": {
                            "de": "USA",
                            "en": "United States",
                            "es": "Estados Unidos",
                            "fr": "États-Unis",
                            "ja": "アメリカ合衆国",
                            "pt-BR": "Estados Unidos",
                            "ru": "США",
                            "zh-CN": "美国"
                        }
                    }
                }
            }
        ]
    },
    {
        "Database": "dbs/GeoLite2-City.mmdb",
        "Records": [
            {
                "Network": "8.0.17.0/24",
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
                        "geoname_id": 6252001,
                        "iso_code": "US",
                        "names": {
                            "de": "USA",
                            "en": "United States",
                            "es": "Estados Unidos",
                            "fr": "États-Unis",
                            "ja": "アメリカ合衆国",
                            "pt-BR": "Estados Unidos",
                            "ru": "США",
                            "zh-CN": "美国"
                        }
                    },
                    "location": {
                        "accuracy_radius": 1000,
                        "latitude": 37.751,
                        "longitude": -97.822,
                        "time_zone": "America/Chicago"
                    },
                    "registered_country": {
                        "geoname_id": 6252001,
                        "iso_code": "US",
                        "names": {
                            "de": "USA",
                            "en": "United States",
                            "es": "Estados Unidos",
                            "fr": "États-Unis",
                            "ja": "アメリカ合衆国",
                            "pt-BR": "Estados Unidos",
                            "ru": "США",
                            "zh-CN": "美国"
                        }
                    }
                }
            }
        ]
    }
]
```
