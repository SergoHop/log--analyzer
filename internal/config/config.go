package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	AppPort string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() *Config{
	_ = godotenv.Load()

	cfg := &Config{
		AppPort: getEnv("APP_PORT", ":8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_PORT", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASSWORD", ""),
		DBName:  getEnv("DB_NAME", "postgres"),
	}

	return cfg
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort,
	)
}

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
