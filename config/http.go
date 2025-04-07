package config

type HTTP struct {
	Port           string `yaml:"port"`
	UsePreforkMode bool   `yaml:"use_prefork_mode"`
}
