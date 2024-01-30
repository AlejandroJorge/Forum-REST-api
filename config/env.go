package config

import (
	"fmt"
	"path"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	currentDir := GetWorkingDir()
	filepath := path.Join(currentDir, ".env")
	err := godotenv.Load(filepath)
	if err != nil {
		fmt.Println("[WARNING] Environment variables not loaded from .env file")
	}
}
