package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	Token    string
}

func ParseEnvVariables() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../deploy/.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

}
