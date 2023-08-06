# Changelog

## 2023-08-06

### Added

- `--target` flag in `chglog changelog` command to specify a different file for the changelog.

### Changed

- Updated `changelogEntry` function to take in two more parameters: `append` and `changelogFile`.
- In `main.go`, added `appendChangelog` and `changelogFile` variables.
- Updated `cli` interface to include `--append` and `--target` flags.
- Constants in `config.go` are now explicitly typed as strings.
- Clarified the instructions in `llm.go` for the formatted markdown list.

### Fixed

- Fixed the changelog entry generation in `actions.go` to conditionally append the message to the specified changelog file.
