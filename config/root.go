package config

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Network struct {
		Port string
	}

	Authenticator struct {
		Secret   string
		Account  string
		Issuer   string
		FileName string
	}

	DB struct {
		Database string
		URL      string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
