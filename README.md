## mmdbinspect

`mmdbinspect` - look up records for one or more IPs/networks in one or more .mmdb databases

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
mmdbinspect [-include-aliased-networks] [-include-build-time] [-include-networks-without-data] [-jsonl] -db path/to/db [IP|network] [IP|network]...
  -db value            Path to an mmdb file. You may pass this arg more than once.
                       This may also be a glob pattern matching one or more MMDB files.
  -include-aliased-networks
                       Include aliased networks (e.g. 6to4, Teredo). This option may
                       cause IPv4 networks to be listed more than once via aliases.
  -include-build-time  Include the build time of the database in the output.
  -include-networks-without-data
                       Include networks that have no data in the database.
                       The "record" will be null for these.
  -jsonl               Output as JSONL instead of YAML.
  [IP|network]         An IP address, or network in CIDR notation. Can be
                       specified multiple times.
```

## Description

Any IPs specified will be treated as their single-host network counterparts (e.g. 1.2.3.4 => 1.2.3.4/32).

`mmdbinspect` will look up each IP/network in each database specified. For each IP/network looked up in a database, the program will select all records for networks which are contained within the looked up IP/network. If no records for contained networks are found in the datafile, the program will select the record that is contained by the looked up IP/network. If no such records are found, none are selected.

The program outputs the selected records in YAML format by default (use `-jsonl` for JSONL format). Each output item corresponds to a single IP/network being looked up in a single DB. Each record contains the following keys: `database_path`, `requested_lookup`, `network`, and `record`. This format allows for efficient streaming of large lookups and makes the key naming more consistent.

## Beta Release

This software is in beta. No guarantees are made, including relating to interface stability. Comments or suggestions for improvements are welcome on our [GitHub issue tracker](https://github.com/maxmind/mmdbinspect/issues).

## Installation

### Installing via binary release

[Release binaries](https://github.com/maxmind/mmdbinspect/releases) have
been made available for several popular platforms. Simply download the
binary for your platform and run it.

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

The easiest way is via `go install`:

```bash
$ go install github.com/maxmind/mmdbinspect/v2/cmd/mmdbinspect@latest
```

This installs `mmdbinspect` to `$GOPATH/bin/mmdbinspect`.

## Examples

<details>
    <summary>A simple lookup (one IP/network, one DB)</summary>

```bash
$ mmdbinspect -db GeoIP2-Country.mmdb 152.216.7.110
database_path: GeoIP2-Country.mmdb
requested_lookup: 152.216.7.110
network: 152.208.0.0/12
record:
  continent:
    code: NA
    geoname_id: 6255149
    names:
      de: Nordamerika
      en: North America
      es: Norteamérica
      fr: Amérique du Nord
      ja: 北アメリカ
      pt-BR: América do Norte
      ru: Северная Америка
      zh-CN: 北美洲
  country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
  registered_country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
```
</details>

<details>
    <summary>Look up one IP/network in multiple databases</summary>

```bash
$ mmdbinspect -db GeoIP2-Country.mmdb -db GeoIP2-City.mmdb 152.216.7.110
database_path: GeoIP2-Country.mmdb
requested_lookup: 152.216.7.110
network: 152.208.0.0/12
record:
  continent:
    code: NA
    geoname_id: 6255149
    names:
      de: Nordamerika
      en: North America
      es: Norteamérica
      fr: Amérique du Nord
      ja: 北アメリカ
      pt-BR: América do Norte
      ru: Северная Америка
      zh-CN: 北美洲
  country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
  registered_country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
---
database_path: GeoIP2-City.mmdb
requested_lookup: 152.216.7.110
network: 152.216.4.0/22
record:
  continent:
    code: NA
    geoname_id: 6255149
    names:
      de: Nordamerika
      en: North America
      es: Norteamérica
      fr: Amérique du Nord
      ja: 北アメリカ
      pt-BR: América do Norte
      ru: Северная Америка
      zh-CN: 北美洲
  country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
  registered_country:
    geoname_id: 6252001
    iso_code: US
    names:
      de: USA
      en: United States
      es: Estados Unidos
      fr: États Unis
      ja: アメリカ
      pt-BR: EUA
      ru: США
      zh-CN: 美国
```
</details>

<details>
    <summary>Look up multiple IPs/networks in a single database</summary>

```bash
$ mmdbinspect -db GeoIP2-Connection-Type.mmdb 152.216.7.110/20 2001:0:98d8::/64
database_path: GeoIP2-Connection-Type.mmdb
requested_lookup: 152.216.7.110/20
network: 152.216.0.0/19
record:
  connection_type: Corporate
---
database_path: GeoIP2-Connection-Type.mmdb
requested_lookup: 2001:0:98d8::/64
network: 2001:0:98d8::/51
record:
  connection_type: Corporate
```
</details>

<details>
    <summary>Look up multiple IPs/networks in multiple databases</summary>

```bash
$ mmdbinspect -db GeoLite2-ASN.mmdb -db GeoIP2-Connection-Type.mmdb 152.216.7.110/20 2001:0:98d8::/64
database_path: GeoIP/GeoLite2-ASN.mmdb
requested_lookup: 152.216.7.110/20
network: 152.216.0.0/19
record:
  autonomous_system_number: 30313
  autonomous_system_organization: IRS
---
database_path: GeoIP/GeoLite2-ASN.mmdb
requested_lookup: 2001:0:98d8::/64
network: 2001:0:98d8::/51
record:
  autonomous_system_number: 30313
  autonomous_system_organization: IRS
---
database_path: GeoIP2-Connection-Type.mmdb
requested_lookup: 152.216.7.110/20
network: 152.216.0.0/19
record:
  connection_type: Cable/DSL
---
database_path: GeoIP2-Connection-Type.mmdb
requested_lookup: 2001:0:98d8::/64
network: 2001:0:98d8::/51
record:
  connection_type: Cable/DSL
```
</details>

<details>
    <summary>Using glob patterns to match multiple database files</summary>

```bash
$ mmdbinspect -db "GeoIP2-*.mmdb" 152.216.7.110
database_path: GeoIP2-Country.mmdb
requested_lookup: 152.216.7.110
network: 152.208.0.0/12
record:
  continent:
    code: NA
    geoname_id: 6255149
    names:
      de: Nordamerika
      en: North America
      # ... more names
  country:
    geoname_id: 6252001
    iso_code: US
    # ... more country data
---
database_path: GeoIP2-City.mmdb
requested_lookup: 152.216.7.110
network: 152.216.4.0/22
record:
  # ... city data
```
</details>

<details>
    <summary>Look up a file of IPs/networks using the <code>xargs</code> utility</summary>

```bash
$ cat list.txt
152.216.7.32/27
2610:30::/64
$ cat list.txt | xargs mmdbinspect -db GeoIP2-ISP.mmdb
database_path: GeoIP2-ISP.mmdb
requested_lookup: 152.216.7.32/27
network: 152.216.0.0/19
record:
  autonomous_system_number: 30313
  autonomous_system_organization: IRS
  isp: Internal Revenue Service
  organization: Internal Revenue Service
---
database_path: GeoIP/GeoIP2-ISP.mmdb
requested_lookup: 2610:30::/64
network: 2610:30::/32
record:
  autonomous_system_number: 30313
  autonomous_system_organization: IRS
  isp: Internal Revenue Service
  organization: Internal Revenue Service
```
</details>

<details>
<summary>Processing the output with the <code>-jsonl</code> flag and the <code>jq</code> utility</summary>

Print out the `isp` field from each result found:
```bash
$ mmdbinspect -jsonl -db GeoIP2-ISP.mmdb 152.216.7.32/27 | jq -r '.record.isp'
Internal Revenue Service
```

Print out the `isp` field from each result found in a specific format using string addition:
```bash
$ mmdbinspect -jsonl -db GeoIP2-ISP.mmdb 152.216.7.32/27 | jq -r '.record | "isp=" + .isp'
isp=Internal Revenue Service
```

Print out the `city` and `country` names from each record using string addition:
```bash
$ mmdbinspect -jsonl -db GeoIP2-City.mmdb 2610:30::/64 | jq -r '.record | .city.names.en + ", " + .country.names.en'
Martinsburg, United States
```

Print out the `city` and `country` names from each record using array construction and `join`:
```bash
$ mmdbinspect -jsonl -db GeoIP2-City.mmdb 2610:30::/64 | jq -r '.record | [.city.names.en, .country.names.en] | join(", ")'
Martinsburg, United States
```

Get the AS number for an IP:
```bash
$ mmdbinspect -jsonl -db GeoLite2-ASN.mmdb 152.216.7.49 | jq -r '.record.autonomous_system_number'
30313
```

Create a CSV file with network and country code for all networks with data:
```bash
$ echo "network,country" > networks.csv
$ mmdbinspect -jsonl -db GeoIP2-Country.mmdb ::/0 | jq -r '[.network, .record.country.iso_code] | join(",")' >> networks.csv
$ cat networks.csv
network,country
1.1.1.0/24,AU
...
```

When asking `jq` to print a path it can't find, it'll print `null`:
```bash
$ mmdbinspect -jsonl -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.invalid.path'
null
```

When asking `jq` to concatenate or join a path it can't find, it'll leave it blank:
```bash
$ mmdbinspect -jsonl -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.record | .city.names.en + ", " + .country.names.en'
, United States
$ mmdbinspect -jsonl -db GeoIP2-City.mmdb 152.216.7.49 | jq -r '.record | [.city.names.en, .country.names.en] | join(", ")'
, United States
```
</details>

<details>
<summary>Using the `-include-*` flags for additional information</summary>

Include build time information:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb -include-build-time 152.216.7.110
database_path: GeoIP2-City.mmdb
build_time: 2023-01-15T12:34:56Z
requested_lookup: 152.216.7.110
network: 152.216.4.0/22
record:
  # ... city data
```

Include networks without data:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb -include-networks-without-data 192.0.2.1
database_path: GeoIP2-City.mmdb
requested_lookup: 192.0.2.1
network: 192.0.2.0/24
```

Include aliased networks:
```bash
$ mmdbinspect -db GeoIP2-City.mmdb -include-aliased-networks ::/0
# ... All IPs in the database, including all aliased networks.
```
</details>

## Bug Reports

Please report bugs by filing an issue with our GitHub issue tracker at [https://github.com/maxmind/mmdbinspect/issues](https://github.com/maxmind/mmdbinspect/issues).

## Copyright and License

This software is Copyright (c) 2019 - 2025 by MaxMind, Inc.

This is free software, licensed under the [Apache License, Version 2.0](LICENSE-APACHE) or the [MIT License](LICENSE-MIT), at your option.
