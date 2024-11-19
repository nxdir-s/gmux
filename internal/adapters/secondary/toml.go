package secondary

import (
	"os"

	"github.com/nxdir-s/gomux/internal/core/valobj"
	"github.com/pelletier/go-toml/v2"
)

const (
	ConfigFile string = "config.toml"
)

type ErrReadCfg struct {
	err error
}

func (e *ErrReadCfg) Error() string {
	return "failed to read " + ConfigFile + ": " + e.err.Error()
}

type ErrUnmarshalToml struct {
	err error
}

func (e *ErrUnmarshalToml) Error() string {
	return "failed to unmarshal " + ConfigFile + ": " + e.err.Error()
}

type Config struct {
	Session     string `toml:"session"`
	Project     string `toml:"project"`
	ServerCmd   string `toml:"server_cmd"`
	DockerCmd   string `toml:"docker_cmd"`
	DatabaseCmd string `toml:"database_cmd"`
}

type TomlAdapter struct {
	cfg *Config
}

func NewTomlAdapter() (*TomlAdapter, error) {
	return &TomlAdapter{}, nil
}

func (a *TomlAdapter) LoadConfig() (*valobj.Config, error) {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, &ErrReadCfg{err}
	}

	err = toml.Unmarshal(data, a.cfg)
	if err != nil {
		return nil, &ErrUnmarshalToml{err}
	}

	return &valobj.Config{
		Session:     a.cfg.Session,
		ProjectDir:  a.cfg.Project,
		ServerCmd:   a.cfg.ServerCmd,
		DockerCmd:   a.cfg.DockerCmd,
		DatabaseCmd: a.cfg.DatabaseCmd,
	}, nil
}
