package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Specification holds environment variable name.
	Specification struct {
		DBHost          string
		DBName          string
		DBDevicesCol    string
		DBUsersCol      string
		DBBorrowingsCol string
		APIPort         string
	}
)

// Spec retrieves the value of the environment variable named by the key.
func Spec() *Specification {
	godotenv.Load("/api/.env")

	s := Specification{
		DBHost:          os.Getenv("DB_HOST"),
		DBName:          os.Getenv("DB_NAME"),
		DBDevicesCol:    os.Getenv("DB_DEVICES_COL"),
		DBUsersCol:      os.Getenv("DB_USERS_COL"),
		DBBorrowingsCol: os.Getenv("DB_BORROWINGS_COL"),
		APIPort:         os.Getenv("API_PORT"),
	}

	return &s
}
