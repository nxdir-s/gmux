package config

const FileName string = ".gomux.toml"

type Window struct {
	Name string
	Cmd  []string
}
