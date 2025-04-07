package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB      DB      `yaml:"db"`
	HTTP    HTTP    `yaml:"http"`
	Swagger Swagger `yaml:"swagger"`
}

func NewConfig() (*Config, error) {
	var err error
	cfg := &Config{}
	if err = cleanenv.ReadConfig("./config.yaml", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	log.Printf("%#v\n", cfg)
	return cfg, nil
}
