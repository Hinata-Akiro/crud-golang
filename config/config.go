
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	JWT_SECRET  string 
}

func Config() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}
	appConfig := &AppConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}
	log.Println("DB_HOST: ", appConfig.DB_HOST)

	return appConfig
}