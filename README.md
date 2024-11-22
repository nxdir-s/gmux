# GoMux

GoMux is a command-line tool that helps with setting up tmux for a project

## Installation

Install GoMux with the following

```shell
$ go install github.com/nxdir-s/gomux/cmd/gomux@latest
```

## Usage

GoMux is intended to be used to setup tmux windows for a project. While in a project directory
run the following to automatically setup tmux

```shell
$ gomux
```

### Configuration

GoMux can be configured using a toml file named `.gomux.toml`. Tmux windows can be configured by adding `windows` sub-tables. A window
requires a `name` and a `cmd` that will be executed

> [TOML documentation](https://toml.io/en/v1.0.0)

#### Example Config

```toml
title = 'Example GoMux Config'

session = 'SessionName'
start_index = 0

[windows]

[windows.editor]
name = 'editor'
cmd = ['vim .']

[windows.docker]
name = 'docker'
cmd = ['docker compose up']

[windows.database]
name = 'database'
cmd = ['psql']

[windows.server]
name = 'server'
cmd = ['go run cmd/server/main.go']
```
