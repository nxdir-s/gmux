# gmux

`gmux` is a command-line tool that helps setting up tmux for a project

## Installation

GoMux can be installed with the following command

```bash
$ go install github.com/nxdir-s/gmux/cmd/gmux@latest
```

## Usage

`gmux` is intended to be used to setup tmux windows for a project. While in a project directory
run the following to automatically setup tmux

```bash
$ gmux
```

### Configuration

`gmux` can be configured using a toml file named `.gmux.toml`. Tmux windows can be configured by adding `windows` sub-tables. A window
requires a `name` and a `cmd` that will be executed

> #### [Toml docs](https://toml.io/en/v1.0.0)

#### Example Config

```toml
title = 'example gmux config'

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
