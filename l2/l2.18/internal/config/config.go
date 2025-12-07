package config

import (
	"flag"
	"os"
)

type Config struct {
	Port string
}

func Load() *Config {
	var res Config
	flag.StringVar(&res.Port, "port", ":8080", "Server port")
	flag.Parse()

	if envPort := os.Getenv("HTTP_PORT"); envPort != "" {
		res.Port = envPort
	}

	return &res
}
