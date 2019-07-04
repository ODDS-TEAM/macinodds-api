package config

import (
	"os"

	"github.com/joho/godotenv"
	"gitlab.odds.team/internship/macinodds-api/model"
)

// Config retrieves the value of the environment variable named by the key.
func Config() *model.Config {
	godotenv.Load()

	config := model.Config{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_MAC_NAME"),
		os.Getenv("DB_MAC_COL"),
		os.Getenv("API_PORT"),
	}

	return &config
}
