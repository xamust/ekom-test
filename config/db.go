package config

import "time"

type DB struct {
	Mongo struct {
		URL            string        `yaml:"url"`
		Port           string        `yaml:"port"`
		MaxPoolSize    uint64        `yaml:"max_pool_size"`
		ConnectTimeout time.Duration `yaml:"connect_timeout"`
		DBName         string        `yaml:"db_name"`
		CollectionName string        `yaml:"collection_name"`
		Username       string        `yaml:"username"`
		Password       string        `yaml:"password"`
	} `yaml:"mongo"`
}
