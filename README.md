# Rules Lint

This cli tool provides comprehensive analysis over declarative rules files.

### How to use

Simply run the following command:

```shell
rules-lint check --config path/to/your/config.yml
```

Below you can check the available commands:

| Command | Description | Example |
|---------|-------------|---------|
| `check` | Lint rule files | `rules-lint check --config config.yaml` |
| `help` | Show help | `rules-lint help` |
| `version` | Show version | `rules-lint version` |


### Config file reference

The config specification is the latest and recommended version of the config file format. It helps you define a configuration file which is used to configure your linting rules.

#### Directories top-level element

The top-level **directories** property is defined by the config specification as the root folder where to look for rules files.

#### Rules top-level element

The top-level rules property is where users specify which lint rules to run.

Example `config.yml`:

```yml
directories:
  - ./rules

rules:
  checkTemplateVars: true
  checkUnusedContextKeys: true
  checkAsyncIncongruence: true
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
