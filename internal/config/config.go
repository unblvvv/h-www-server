package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database  DatabaseConfig
	DSN       string
	JWTSecret string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	databaseConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", ""),
		Name:     getEnv("POSTGRES_DB", "biteway"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.Name, databaseConfig.SSLMode)

	cfg := &Config{
		Database:  databaseConfig,
		DSN:       dsn,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	return cfg, nil

}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
