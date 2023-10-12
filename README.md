# YAML-ENV-RUNNER

The yaml-env-runner is a simple program that reads environment variables from a YAML file and runs a specified program
using those environment variables.

## Usage

```bash
$ yaml-env-runner -f=<path-to-yaml-file> -- <command> <command-args>
```

Where `<path-to-yaml-file>` is the path to the YAML file that contains the environment variables key-value pairs
and `<command> <command-args>` indicates the program to be executed using the environment variables.

Example YAML file:

```yaml
DB_HOST: localhost
DB_PORT: 5432
DB_USER: postgres
DB_PASS: secret
```

To run a Python script using the environment variables in the `env.yaml` file, execute:

```bash
$ yaml-env-runner -f=env.yaml -- python script.py
```

## Features

1. Simplifies the management of environment variables for different projects or environments.

2. Loads environment variables from a YAML file before executing the specified program

3. Supports execution of any program with arguments

## Requirements

- Go 1.16 or above

## Installation

- Clone the repository

```bash
$ git clone https://github.com/yourusername/yaml-env-runner.git
```

- Build the program

```bash
$ go build -o yaml-env-runner main.go
```

This will generate an executable named `yaml-env-runner`.

## License

This project is licensed under the MIT License.