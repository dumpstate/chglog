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
