# go-in-practice

Notes and code examples created whilst following the [Go In Practice](https://www.manning.com/books/go-in-practice) book.

## Techniques

### CLI

- [flag_cli.go](/cli/1/flag_cli.go) - Using Go's built-in flag system (based on Plan 9)

#### GNU/Unix-style command-line arguments

- [flag_based_cli.go](/cli/2/flag_based_cli.go) - drop in replacement
- [go_flags_cli.go](/cli/3/go_flags_cli.go) - different api

#### Avoiding CLI boilerplate code

- [hello_world.go](/cli/4/hello_world.go) - basic example
- [count_cli.go](/cli/5/count_cli.go) - cli that supports multiple commands and subcommands 

### Config

- [json_config.go](/config/1/json_config.go) - loading config from a json file
- [yaml_config.go](/config/2/yaml_config.go) - loading config from a yaml file
- [ini_config.go](/config/3/ini_config.go) - loading config from an ini file
- [env_config.go](/config/4/env_config.go) - retreive config from the environment
