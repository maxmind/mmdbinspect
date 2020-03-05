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

The following examples were created using our test databases which can be found in `test/data` in this repository.

<details>
    <summary>A simple lookup (one IP/network, one DB)</summary>

```bash
$ mmdbinspect -db GeoIP2-Anonymous-IP-Test.mmdb 71.160.223.3
[
    {
        "Database": "GeoIP2-Anonymous-IP-Test.mmdb",
        "Records": [
            {
                "Network": "71.160.223.3/24",
                "Record": {
                    "is_anonymous": true,
                    "is_hosting_provider": true
                }
            }
        ],
        "Lookup": "71.160.223.3"
    }
]
```
</details>

<details>
    <summary>Look up one IP/network in multiple databases</summary>

```bash
$ mmdbinspect -db GeoIP2-Country-Test.mmdb -db GeoIP2-City-Test.mmdb 202.196.224.4
[
    {
        "Database": "GeoIP2-Country-Test.mmdb",
        "Records": [
            {
                "Network": "202.196.224.4/20",
                "Record": {
                    "continent": {
                        "code": "AS",
                        "geoname_id": 6255147,
                        "names": {
                            "de": "Asien",
                            "en": "Asia",
                            "es": "Asia",
                            "fr": "Asie",
                            "ja": "アジア",
                            "pt-BR": "Ásia",
                            "ru": "Азия",
                            "zh-CN": "亚洲"
                        }
                    },
                    "country": {
                        "geoname_id": 1694008,
                        "iso_code": "PH",
                        "names": {
                            "de": "Philippinen",
                            "en": "Philippines",
                            "es": "Filipinas",
                            "fr": "Philippines",
                            "ja": "フィリピン共和国",
                            "pt-BR": "Filipinas",
                            "ru": "Филиппины",
                            "zh-CN": "菲律宾"
                        }
                    },
                    "registered_country": {
                        "geoname_id": 1694008,
                        "iso_code": "PH",
                        "names": {
                            "de": "Philippinen",
                            "en": "Philippines",
                            "es": "Filipinas",
                            "fr": "Philippines",
                            "ja": "フィリピン共和国",
                            "pt-BR": "Filipinas",
                            "ru": "Филиппины",
                            "zh-CN": "菲律宾"
                        }
                    },
                    "represented_country": {
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
                        },
                        "type": "military"
                    }
                }
            }
        ],
        "Lookup": "202.196.224.4"
    },
    {
        "Database": "GeoIP2-City-Test.mmdb",
        "Records": [
            {
                "Network": "202.196.224.4/20",
                "Record": {
                    "continent": {
                        "code": "AS",
                        "geoname_id": 6255147,
                        "names": {
                            "de": "Asien",
                            "en": "Asia",
                            "es": "Asia",
                            "fr": "Asie",
                            "ja": "アジア",
                            "pt-BR": "Ásia",
                            "ru": "Азия",
                            "zh-CN": "亚洲"
                        }
                    },
                    "country": {
                        "geoname_id": 1694008,
                        "iso_code": "PH",
                        "names": {
                            "de": "Philippinen",
                            "en": "Philippines",
                            "es": "Filipinas",
                            "fr": "Philippines",
                            "ja": "フィリピン共和国",
                            "pt-BR": "Filipinas",
                            "ru": "Филиппины",
                            "zh-CN": "菲律宾"
                        }
                    },
                    "location": {
                        "accuracy_radius": 121,
                        "latitude": 13,
                        "longitude": 122,
                        "time_zone": "Asia/Manila"
                    },
                    "postal": {
                        "code": "34021"
                    },
                    "registered_country": {
                        "geoname_id": 1694008,
                        "iso_code": "PH",
                        "names": {
                            "de": "Philippinen",
                            "en": "Philippines",
                            "es": "Filipinas",
                            "fr": "Philippines",
                            "ja": "フィリピン共和国",
                            "pt-BR": "Filipinas",
                            "ru": "Филиппины",
                            "zh-CN": "菲律宾"
                        }
                    },
                    "represented_country": {
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
                        },
                        "type": "military"
                    }
                }
            }
        ],
        "Lookup": "202.196.224.4"
    }
]
```
</details>

<details>
    <summary>Look up multiple IPs/networks in a single database</summary>

```bash
$ mmdbinspect -db GeoIP2-Connection-Type-Test.mmdb 1.0.9.16/28 2003::ff00:3a4c/128
[
    {
        "Database": "GeoIP2-Connection-Type-Test.mmdb",
        "Records": [
            {
                "Network": "1.0.9.16/21",
                "Record": {
                    "connection_type": "Dialup"
                }
            }
        ],
        "Lookup": "1.0.9.16/28"
    },
    {
        "Database": "GeoIP2-Connection-Type-Test.mmdb",
        "Records": [
            {
                "Network": "2003::ff00:3a4c/24",
                "Record": {
                    "connection_type": "Cable/DSL"
                }
            }
        ],
        "Lookup": "2003::ff00:3a4c/128"
    }
]
```
</details>

<details>
    <summary>Look up multiple IPs/networks in multiple databases</summary>

```bash
$ mmdbinspect -db GeoIP2-Static-IP-Score-Test.mmdb -db GeoIP2-User-Count-Test.mmdb 1.2.3.5 1.2.3.66
[
    {
        "Database": "GeoIP2-Static-IP-Score-Test.mmdb",
        "Records": [
            {
                "Network": "1.2.3.5/32",
                "Record": {
                    "score": 0.06
                }
            }
        ],
        "Lookup": "1.2.3.5"
    },
    {
        "Database": "GeoIP2-Static-IP-Score-Test.mmdb",
        "Records": [
            {
                "Network": "1.2.3.66/26",
                "Record": {
                    "score": 0.12
                }
            }
        ],
        "Lookup": "1.2.3.66"
    },
    {
        "Database": "GeoIP2-User-Count-Test.mmdb",
        "Records": [
            {
                "Network": "1.2.3.5/32",
                "Record": {
                    "ipv4_24": 4,
                    "ipv4_32": 1
                }
            }
        ],
        "Lookup": "1.2.3.5"
    },
    {
        "Database": "GeoIP2-User-Count-Test.mmdb",
        "Records": [
            {
                "Network": "1.2.3.66/26",
                "Record": {
                    "ipv4_24": 4,
                    "ipv4_32": 0
                }
            }
        ],
        "Lookup": "1.2.3.66"
    }
]
```
</details>

<details>
    <summary>Look up a file of IPs/networks using the <code>xargs</code> utility</summary>

```bash
$ cat list.txt
5.83.124.0/20
216.160.83.0/27
$ cat list.txt | xargs mmdbinspect -db GeoIP2-DensityIncome-Test.mmdb
[
    {
        "Database": "GeoIP2-DensityIncome-Test.mmdb",
        "Records": [
            {
                "Network": "5.83.124.0/22",
                "Record": {
                    "average_income": 32323,
                    "population_density": 1232
                }
            }
        ],
        "Lookup": "5.83.124.0/20"
    },
    {
        "Database": "GeoIP2-DensityIncome-Test.mmdb",
        "Records": [
            {
                "Network": "216.160.83.0/24",
                "Record": {
                    "average_income": 24626,
                    "population_density": 1341
                }
            }
        ],
        "Lookup": "216.160.83.0/27"
    }
]
```
</details>

<details>
<summary>Tame the output with the <code>jq</code> utility</summary>

```bash
$ mmdbinspect -db GeoIP2-ISP-Test.mmdb 5.145.96.0 | jq '.[] | .Records[].Record.isp'
"Finecom"
```
</details>

## Bug Reports

Please report bugs by filing an issue with our GitHub issue tracker at [https://github.com/maxmind/mmdbinspect/issues](https://github.com/maxmind/mmdbinspect/issues).

## Copyright and License

This software is Copyright (c) 2019 - 2020 by MaxMind, Inc.

This is free software, licensed under the [Apache License, Version 2.0](LICENSE-APACHE) or the [MIT License](LICENSE-MIT), at your option.
