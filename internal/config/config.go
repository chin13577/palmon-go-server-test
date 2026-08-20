package config

import "os"

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	return &Config{
		Port:       env("PORT", "4001"),
		DBHost:     env("DB_HOST", ""),
		DBPort:     env("DB_PORT", ""),
		DBUser:     env("DB_USERNAME", ""),
		DBPassword: env("DB_PASSWORD", ""),
		DBName:     env("DB_DATABASE", ""),
	}
}

func env(key string, fallback string) string {
	result := os.Getenv(key)
	if result == "" {
		return fallback
	} else {
		return result
	}
}
