package env

import (
	"ctfbot/internal/values"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var cfg Config

func Load() {
	godotenv.Load()

	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("failed to load env: %v", err)
	}

	switch cfg.Mode {
	case values.Dev:
		logrus.SetLevel(logrus.DebugLevel)
	case values.Preprod:
		logrus.SetLevel(logrus.DebugLevel)
	case values.Prod:
		logrus.SetLevel(logrus.InfoLevel)
	case values.Unset:
		logrus.Fatalf("MODE is not set, be sure to have a .env file or set the environment variables")
	default:
		logrus.Fatalf("MODE is not set, be sure to have a .env file or set the environment variables")
	}

	logrus.Infof("Environment loaded: %s", cfg.Mode)
	logrus.Debugf("Environment: %+v", cfg)
}
