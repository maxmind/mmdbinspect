# Releasing

* Install `goreleaser`. Refer to its docs.
* Set a `GITHUB_TOKEN` environment variable. Refer to `goreleaser` docs for
  information.
* Update `CHANGELOG.md`.
  * Mention recent changes.
  * Set a version if there is not one.
  * Set a release date.
* Commit `CHANGELOG.md`.
* Tag the release: `git tag -a v1.2.3 -m 'Tag v1.2.3'`.
* Push the tag: `git push origin v1.2.3`.
* Run `goreleaser`.
