package config

import (
	"fmt"
	"os"
)

type config struct {
	Port       string
	DBUser     string
	DBPasswd   string
	PublicHost string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() config {

	return config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPasswd:   getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
