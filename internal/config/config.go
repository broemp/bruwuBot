package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token  string `json:"bot_token"`
	Prefix string `json:"prefix"`
}

func ParseConfig(fileName string) (c *Config, err error) {

	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(Config)
	err = json.NewDecoder(f).Decode(c)

	return
}
