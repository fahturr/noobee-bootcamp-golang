package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App  App  `yaml:"app"`
	Mail Mail `yaml:"mail"`
}

type App struct {
	Port string `yaml:"port"`
}

type Mail struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

var Cfg *Config

func LoadConfig(filename string) (err error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	cfg := Config{}

	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg
	return
}
