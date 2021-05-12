package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigForYaml() (*Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("loadConfigForYaml os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
