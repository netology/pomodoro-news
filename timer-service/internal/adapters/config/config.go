package config

import (
	"time"
)

func MustLoad(configPath string) *Config {
	return &Config{
		Port:        8080,
		ServiceName: "timer-service",
		DSN:         "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		Timeout:     30 * time.Second,
		LogLevel:    "INFO",
		LogFormat:   "JSON",
	}
}

type Config struct {
	Port        int
	ServiceName string
	DSN         string
	Timeout     time.Duration
	LogLevel    string
	LogFormat   string
}
