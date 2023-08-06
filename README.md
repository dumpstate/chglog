# chglog

Maintain changelog with the power of generative AI.

## Usage

### Initialise Config

```sh
chglog init
```

The `chglog` will ask for required configuration parameters and store in the config directory.

Subsequent call will override the current config.

Note, to configure the details, one should edit the configuration file manually.

### Generate Commit Message

Call from the git repository:

```sh
chglog commit
```

`chglog` reads the branch name, takes the `git diff` and asks LLM for a single line commit message.

### Generate Changelog Entry

Generates changelog entry and appends to `CHANGELOG.md` file in the root of the project.

```sh
chglog changelog [--append] [--target CHANGELOG.md]
```

`chglog` takes a diff of all unstaged changes and asks LLM for a changelog entry.

By default, the changelog entry is just printed to stdout, with `--append` flag, it is appended
to the `CHANGELOG.md` file. The `--target` flag allows to specify a different file.
