## mmdbinspect

`mmdbinspect` - look up records for one or more IPs/networks in one or more .mmdb databases

[![Build Status](https://github.com/maxmind/mmdbinspect/workflows/Build/badge.svg)](https://github.com/maxmind/mmdbinspect/actions?query=workflow%3A%22Build%22)
[![Coverage Status](https://coveralls.io/repos/github/maxmind/mmdbinspect/badge.svg)](https://coveralls.io/github/maxmind/mmdbinspect)

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Description](#description)
* [Beta Release](#beta-release)
* [Installation](#installation)
    * [Installing via binary release](#installing-via-binary-release)
    * [Installing on Linux via the tarball](#installing-on-linux-via-the-tarball)
    * [Installing on Ubuntu or Debian via the .deb](#installing-on-ubuntu-or-debian-via-the-deb)
    * [Installing on RedHat or CentOS via the rpm](#installing-on-redhat-or-centos-via-the-rpm)
    * [Installing on macOS (darwin) via the tarball](#installing-on-macos-darwin-via-the-tarball)
    * [Installing on Windows](#installing-on-windows)
    * [Installing from source or Git](#installing-from-source-or-git)
* [Examples](#examples)
* [Bug Reports](#bug-reports)
* [Copyright and License](#copyright-and-license)

<!-- vim-markdown-toc -->

## Usage

```bash
mmdbinspect -db <path/to/your.mmdb> <IP|network>
  db            Path to a MMDB file. Can be specified multiple times.
  <IP|network>  An IP address, or network in CIDR notation. Can be
                specified multiple times.
```

## Description

Any IPs specified will be treated as their single-host network counterparts (e.g. 1.2.3.4 => 1.2.3.4/32).

`mmdbinspect` will look up each IP/network in each database specified. For each IP/network looked up in a database, the program will select all records for networks which are contained within the looked up IP/network. If no records for contained networks are found in the datafile, the program will select the record that is contained by the looked up IP/network. If no such records are found, none are selected.

The program outputs the selected records as a JSON array, with each item in the array corresponding to a single IP/network being looked up in a single DB. The `Database` and `Lookup` keys are added to each item to help correlate which set of records resulted from looking up which IP/network in which database.

## Beta Release

This software is in beta. No guarantees are made, including relating to interface stability. Comments or suggestions for improvements are welcome on our [GitHub issue tracker](https://github.com/maxmind/mmdbinspect/issues).

## Installation

### Installing via binary release

[Release binaries](https://github.com/maxmind/mmdbinspect/releases) have been made available for several popular platorms. Simply download the binary for your platform and run it.

### Installing on Linux via the tarball

Download and extract the appropriate tarball for your system. You will end up with a directory named something like `mmdbinspect_0.0.0_linux_amd64` depending on the version and architecture.

Copy `mmdbinspect` to where you want it to live. To install it into `/usr/local/bin/mmdbinspect`, run the equivalent of `sudo cp mmdbinspect_0.0.0_linux_amd64/mmdbinspect /usr/local/bin`.

### Installing on Ubuntu or Debian via the .deb

(N.B. You can also use the tarball.)

Download the appropriate .deb for your system.

Run `dpkg -i path/to/mmdbinspect_0.0.0_linux_amd64.deb` (replacing the version number and architecture as necessary). You will need to be root. For Ubuntu you can prefix the command with sudo. This will install `mmdbinspect` to `/usr/bin/mmdbinspect`.

### Installing on RedHat or CentOS via the rpm

(N.B. You can also use the tarball.)

Run `rpm -i path/to/mmdbinspect_0.0.0_linux_amd64.rpm` (replacing the version number and architecture as necessary). You will need to be root. This will install `mmdbinspect` to `/usr/bin/mmdbinspect`.

### Installing on macOS (darwin) via the tarball

This is the same as installing on Linux via the tarball, except choose a tarball with "darwin" in the name.

### Installing on Windows

Download and extract the appropriate zip for your system. You will end up with a directory named something like mmdbinspect_0.0.0_windows_amd64 depending on the version and architecture.

Copy mmdbinspect.exe to where you want it to live.

### Installing from source or Git

_We aim always to support the current and penultimate major releases of the Go compiler. You can get it at the [Go website](https://golang.org)._

The easiest way is via `go get`:

```bash
$ go get -u github.com/maxmind/mmdbinspect/cmd/mmdbinspect
```

This installs `mmdbinspect` to `$GOPATH/bin/mmdbinspect`.

## Examples

<details>
    <summary>A simple lookup (one IP/network, one DB)</summary>

```bash
$ mmdbinspect -db GeoIP2-Country.mmdb 152.216.7.110
[
    {
        "Database": "GeoIP2-Country.mmdb",
        "Records": [
            {
                "Network": "152.216.7.110/12",
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
        ],
        "Lookup": "152.216.7.110"
    }
]
```
</details>

<details>
    <summary>Look up one IP/network in multiple databases</summary>

```bash
$ mmdbinspect -db GeoIP2-Country.mmdb -db GeoIP2-City.mmdb 152.216.7.110
[
    {
        "Database": "GeoIP2-Country.mmdb",
        "Records": [
            {
                "Network": "152.216.7.110/12",
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
        ],
        "Lookup": "152.216.7.110"
    },
    {
        "Database": "GeoIP2-City.mmdb",
        "Records": [
            {
                "Network": "152.216.7.110/14",
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
        ],
        "Lookup": "152.216.7.110"
    }
]
```
</details>

<details>
    <summary>Look up multiple IPs/networks in a single database</summary>

```bash
$ mmdbinspect -db GeoIP2-Connection-Type.mmdb 152.216.7.110/20 2001:0:98d8::/64
[
    {
        "Database": "GeoIP2-Connection-Type.mmdb",
        "Records": [
            {
                "Network": "152.216.0.0/13",
                "Record": {
                    "connection_type": "Corporate"
                }
            }
        ],
        "Lookup": "152.216.7.110/20"
    },
    {
        "Database": "GeoIP2-Connection-Type.mmdb",
        "Records": [
            {
                "Network": "2001:0:98d8::/45",
                "Record": {
                    "connection_type": "Corporate"
                }
            }
        ],
        "Lookup": "2001:0:98d8::/64"
    }
]
```
</details>

<details>
    <summary>Look up multiple IPs/networks in multiple databases</summary>

```bash
$ mmdbinspect -db GeoIP2-DensityIncome.mmdb -db GeoIP2-User-Count.mmdb 152.216.7.32/27 2610:30::/64
[
    {
        "Database": "GeoIP2-DensityIncome.mmdb",
        "Records": [
            {
                "Network": "152.216.7.32/21",
                "Record": {
                    "average_income": 26483,
                    "population_density": 1265
                }
            }
        ],
        "Lookup": "152.216.7.32/27"
    },
    {
        "Database": "GeoIP2-DensityIncome.mmdb",
        "Records": [
            {
                "Network": "2610:30::/38",
                "Record": {
                    "average_income": 30369,
                    "population_density": 934
                }
            }
        ],
        "Lookup": "2610:30::/64"
    },
    {
        "Database": "GeoIP2-User-Count.mmdb",
        "Records": [
            {
                "Network": "152.216.7.32/27",
                "Record": {
                    "ipv4_24": 6,
                    "ipv4_32": 0
                }
            }
        ],
        "Lookup": "152.216.7.32/27"
    },
    {
        "Database": "GeoIP2-User-Count.mmdb",
        "Records": [
            {
                "Network": "2610:30::/27",
                "Record": {
                    "ipv6_32": 0,
                    "ipv6_48": 0,
                    "ipv6_64": 0
                }
            }
        ],
        "Lookup": "2610:30::/64"
    }
]
```
</details>

<details>
    <summary>Look up a file of IPs/networks using the <code>xargs</code> utility</summary>

```bash
$ cat list.txt
152.216.7.32/27
2610:30::/64
$ cat list.txt | xargs mmdbinspect -db GeoIP2-ISP.mmdb
[
    {
        "Database": "/usr/local/share/GeoIP/GeoIP2-ISP.mmdb",
        "Records": [
            {
                "Network": "152.216.7.32/20",
                "Record": {
                    "autonomous_system_number": 30313,
                    "autonomous_system_organization": "IRS",
                    "isp": "Internal Revenue Service",
                    "organization": "Internal Revenue Service"
                }
            }
        ],
        "Lookup": "152.216.7.32/27"
    },
    {
        "Database": "/usr/local/share/GeoIP/GeoIP2-ISP.mmdb",
        "Records": [
            {
                "Network": "2610:30::/32",
                "Record": {
                    "autonomous_system_number": 30313,
                    "autonomous_system_organization": "IRS",
                    "isp": "Internal Revenue Service",
                    "organization": "Internal Revenue Service"
                }
            }
        ],
        "Lookup": "2610:30::/64"
    }
]
```
</details>

<details>
<summary>Tame the output with the <code>jq</code> utility</summary>

Print out the `isp` field from each result found:
```bash
$ mmdbinspect -db GeoIP2-ISP.mmdb 152.216.7.32/27 | jq -r '.[].Records[].Record.isp'
Internal Revenue Service
```

Print out the `isp` field from each result found in a specific format using string addition:
```bash
$ mmdbinspect -db GeoIP2-ISP.mmdb 152.216.7.32/27 | jq -r '.[].Records[].Record | "isp=" + .isp'
isp=Internal Revenue Service
```

Print out the `city` and `country` names from each record using string addition:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb 2610:30::/64 | jq -r '.[].Records[].Record | .city.names.en + ", " + .country.names.en'
Martinsburg, United States
```

Print out the `city` and `country` names from each record using array construction and `join`:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb 2610:30::/64 | jq -r '.[].Records[].Record | [.city.names.en, .country.names.en] | join(", ")'
Martinsburg, United States
```

Get the AS number for an IP:
```bash
$ mmdbinspect -db GeoLite2-ASN.mmdb 152.216.7.49 | jq -r '.[].Records[].Record.autonomous_system_number'
30313
```

When asking `jq` to print a path it can't find, it'll print `null`:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.[].invalid.path'
null
```

When asking `jq` to concatenate or join a path it can't find, it'll leave it blank:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.[].Records[].Record | .city.names.en + ", " + .country.names.en'
, United States
$ mmdbinspect -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.[].Records[].Record | [.city.names.en, .country.names.en] | join(", ")'
, United States
```
</details>

## Bug Reports

Please report bugs by filing an issue with our GitHub issue tracker at [https://github.com/maxmind/mmdbinspect/issues](https://github.com/maxmind/mmdbinspect/issues).

## Copyright and License

This software is Copyright (c) 2019 - 2021 by MaxMind, Inc.

This is free software, licensed under the [Apache License, Version 2.0](LICENSE-APACHE) or the [MIT License](LICENSE-MIT), at your option.
