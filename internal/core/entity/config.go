package entity

import "github.com/nxdir-s/gomux/internal/core/entity/config"

type Config struct {
	Session    string
	StartIndex int
	Windows    []config.Window
}
