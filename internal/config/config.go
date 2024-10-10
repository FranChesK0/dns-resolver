package config

import "os"

type Config struct {
	NameServer string
}

func NewConfig() *Config {
	nameServer := os.Getenv("RESOLVE_NAME_SERVER")
	if nameServer == "" {
		nameServer = "77.240.157.30"
	}
	return &Config{NameServer: nameServer}
}
