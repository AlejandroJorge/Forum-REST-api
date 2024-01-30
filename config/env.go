package config

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/AlejandroJorge/forum-rest-api/util"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	currentDir := util.GetWorkingDir()
	filepath := path.Join(currentDir, ".env")
	err := godotenv.Load(filepath)
	if err != nil {
		fmt.Println("[WARNING] Environment variables not loaded from .env file")
	}
}

func getEnvUint(key string) (uint, bool) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}

	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, false
	}

	return uint(parsed), true
}

func getEnvString(key string) (string, bool) {
	return os.LookupEnv(key)
}

func getEnvBytes(key string) ([]byte, bool) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return nil, false
	}

	parsed := []byte(value)
	return parsed, true
}
