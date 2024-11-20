# GoMux

GoMux is a CLI tool that helps with tmux setup for your projects

## Installation

Install GoMux with the following

```shell
$ go install github.com/nxdir-s/gomux/cmd/gomux@latest
```

## Usage

GoMux is intended to be used to setup your tmux windows for a project. While in your project directory
run the following to automatically setup tmux using a config file

```shell
$ gomux
```

### Configuration

GoMux can be configured using a toml file named `.gomux.toml`. Tmux windows can be configured by adding `windows` sub-tables. A window
requires a `name` and a `cmd` that will be executed

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
