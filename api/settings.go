package api

import (
	"github.com/joho/godotenv"
)

func LoadSettings() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
