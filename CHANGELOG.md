# CHANGELOG

## 2.0.0

* Upgrade to `github.com/oschwald/maxminddb-golang/v2`. This is a breaking
  API change, but should not affect the use of the program.
* You may now use a glob for the `-db` argument. If there are multiple
  matches, it will be treated as if multiple `-db` arguments were provided.
  Note that you must quote the parameter when using globs to prevent the
  shell's globbing from interfering. See the [pattern syntax](https://pkg.go.dev/path#Match)


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
