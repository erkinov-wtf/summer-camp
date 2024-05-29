package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Cfg Config

type Config struct {
	App      App
	Database Database
}

type App struct {
	Port     string
	Timezone string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Timezone string
}

func MustLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Cfg = Config{
		App: App{
			Port:     getEnv("PORT", ""),
			Timezone: getEnv("APP_TIMEZONE", "Asia/Tashkent"),
		},
		Database: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "postgres"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Timezone: getEnv("DB_TIMEZONE", "Asia/Tashkent"),
		},
	}

	// Log the loaded configuration for debugging
	log.Printf("Configuration loaded: %+v", Cfg)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Println(value)
	return value
}
