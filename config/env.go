package config

import (
	"path"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	currentDir := GetWorkingDir()
	filepath := path.Join(currentDir, ".env")
	err := godotenv.Load(filepath)
	if err != nil {
		panic(err)
	}
}
