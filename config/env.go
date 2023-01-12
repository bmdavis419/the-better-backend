package config

import (
	"os"

	"github.com/joho/godotenv"
)

// this function will load the .env file if the GO_ENV environment variable is not set
func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
