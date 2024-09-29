package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	ApiUrl      string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	config := Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ApiUrl:      os.Getenv("API_URL"),
	}
	return config
}
