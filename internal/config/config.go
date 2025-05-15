package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() *Config {
	// Load .env file values into OS env vars
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using fallback values where needed")
	}

	return &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "go_cruddb"),
	}
}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}
