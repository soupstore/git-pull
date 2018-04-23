package config

import (
	"github.com/codingconcepts/env"
)

type Config struct {
	DataFolder string `env:"DATA_FOLDER"`
	Repository string `env:"REPO"`
	Branch     string `env:"BRANCH"`
}

func Load() (config *Config, err error) {
	config = &Config{}

	if err = env.Set(config); err != nil {
		return nil, err
	}

	return
}
