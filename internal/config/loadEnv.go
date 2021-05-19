package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientId     string `envconfig:"CLIENT_ID" yaml:"clientId"`
	ClientSecret string `envconfig:"CLIENT_SECRET"  yaml:"clientSecret"`
	RedirectURL  string `envconfig:"REDIRECT_URL" yaml:"RedirectURL"`
	CompanyId    int32  `envconfig:"COMPANY_ID"  yaml:"CompanyId"`
	Filter       struct {
		Items    []int32 `yaml:"Items"`
		Sections []int32 `yaml:"Sections"`
		Tags     []int32 `yaml:"Tags"`
	} `yaml:"Filter"`
}

func LoadConfigForEnv() (*Config, error) {
	var cfg Config
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return &cfg, err
	}

	err = envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println(err.Error())
		return &cfg, err
	}
	return &cfg, err
}
