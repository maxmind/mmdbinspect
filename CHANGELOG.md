# CHANGELOG

## 2.0.0 (2025-11-12)

* **BREAKING CHANGE**: Networks with empty map or empty array data are now
  skipped by default. This makes the output cleaner when working with
  databases that use empty values for networks without meaningful data (e.g.,
  public IPs in some databases). Use the new `-include-empty-values` flag to
  include these networks in the output.
* Added `-include-empty-values` flag to optionally include networks whose data
  is an empty map or empty array.

## 2.0.0-beta.2 (2025-03-11)

* Fixes for the release scripts. No other changes.

## 2.0.0-beta.1 (2025-03-11)

* The default output format is now YAML. This was done to improve the
  readability when using the tool as a standalone utility for doing lookups
  in an MMDB database. Use the `-jsonl` flag to output as JSONL instead.
* When outputting as JSON, we now use JSONL. There is one JSON object per
  line.
* The output format has been flattened. Each record that is output now
  contains the following keys: `database_path`, `requested_lookup`,
  `network`, and `record`. This allows for efficient streaming of large
  lookups, makes the key naming more consistent, and reduces the depth of
  the data structure.
* You may now use a glob for the `-db` argument. If there are multiple
  matches, it will be treated as if multiple `-db` arguments were provided.
  Note that you must quote the parameter when using globs to prevent the
  shell's globbing from interfering. See the [pattern syntax](https://pkg.go.dev/path#Match)
* The following flags were added:
  * `-include-networks-without-data` - include networks without any data in
    the database in the output.
  * `-include-build-time` - include the build time from the database's
    metadata in the output.
* This repo no longer provides a public Go API. It is only intended to be
  used as a CLI program.

## 0.2.0 (2024-01-10)

* Don't escape `&`, `<`, and `>` in JSON output
* Skip aliased IPv6 networks by default
* Build and test with Go 1.21
* Remove deprecated use of ioutil and pkg/errors
* Update dependencies
* Update documentation

## 0.1.1 (2020-02-18)

* Fix release config
* Add release instructions

## 0.1.0 (2020-02-18)

* Initial beta release
