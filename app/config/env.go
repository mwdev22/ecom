package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host string
	Name string
	User string
	Pass string
}

func GetDbCfg() DbConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return DbConfig{
		Host: os.Getenv("DB_HOST"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASSWORD"),
	}
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))
