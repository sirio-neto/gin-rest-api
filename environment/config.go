package config

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort    string `json:"API_PORT"`
	DbPort     string `json:"DATABASE_PORT"`
	DbHost     string `json:"DATABASE_HOST"`
	DbUser     string `json:"DATABASE_USER"`
	DbPassword string `json:"DATABASE_PASSWORD"`
	DbName     string `json:"DATABASE_NAME"`
	DbSslMode  string `json:"DATABASE_SSLMODE"`
}

var Env Config

func InitEnvironmentConfig() {
	if err := godotenv.Load(); err != nil {
		LogInitEnvironmentError(err)
	}

	var readedEnv map[string]string
	readedEnv, err := godotenv.Read()

	if err != nil {
		LogInitEnvironmentError(err)
	}

	parsedJsonEnv, _ := json.Marshal(readedEnv)
	if err := json.Unmarshal(parsedJsonEnv, &Env); err != nil {
		LogInitEnvironmentError(err)
	}
}

func LogInitEnvironmentError(err error) {
	log.Panic("Ocorreu um erro ao iniciar variáveis de ambiente", err)
}
