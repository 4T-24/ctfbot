package app

import (
	"ctfbot/internal/env"
)

type Option func(a *App)

func WithConfig(config *env.Config) Option {
	return func(a *App) {
		a.config = config
	}
}
