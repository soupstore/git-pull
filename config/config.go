package config

import (
	"github.com/codingconcepts/env"
)

type Config struct {
	DataFolder string `env:"DATA_FOLDER" default:"/etc/data"`
	Repository string `env:"REPO"`
	Branch     string `env:"BRANCH" default:"master"`
	SSHPath    string `env:"SSH_PATH" default:"/etc/creds/key"`
}

func Load() (config *Config, err error) {
	config = &Config{}

	if err = env.Set(config); err != nil {
		return nil, err
	}

	return
}
