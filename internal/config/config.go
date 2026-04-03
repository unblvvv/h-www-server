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

	Port string

	R2AccessKey string
	R2SecretKey string
	R2Endpoint  string
	R2Bucket    string
	R2PublicURL string
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

		Port: getEnv("PORT", "8080"),

		R2AccessKey: os.Getenv("R2_ACCESS_KEY"),
		R2SecretKey: os.Getenv("R2_SECRET_KEY"),
		R2Endpoint:  os.Getenv("R2_ENDPOINT"),
		R2Bucket:    os.Getenv("R2_BUCKET"),
		R2PublicURL: os.Getenv("R2_PUBLIC_URL"),
	}

	return cfg, nil

}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
