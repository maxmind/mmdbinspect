# Prereqs

- You must have the [GitHub CLI tool (gh)](https://cli.github.com/) installed,
  in your path, and logged into an account that can push to the repo.
- Your environment also must have `bash`, `git` and `sed` available.

# Releasing

- Review open issues and PRs to see if anything needs to be addressed before
  release.
- Create a branch e.g. `horgh/release` and switch to it.
  - `main` is protected.
- Set the release version and release date in `CHANGELOG.md`. Be sure the
  version follows [Semantic Versionsing](https://semver.org/).
  - Mention recent changes if needed.
- Commit these changes.
- Run `dev-bin/release.sh`.
  - You might need to initialize/update submodules to successfully run tests,
    eg. `git submodule update --init --recursive`.
  - The script pushes the branch. Then it pushes an annotated tag, for example
    `v2.0.1`. The tag message contains the release notes from `CHANGELOG.md`.
- The tag push starts the Release workflow. The authorized releasers receive an
  email to review the pending deployment. If you are an authorized releaser,
  approve the deployment. If you are not, wait for an authorized releaser to do
  so.
- After the approval, GoReleaser creates the GitHub release with the notes from
  the tag. It uploads the binaries and packages, then publishes the release.
- Verify the release on the GitHub Releases page.
- Make a PR and get it merged.
