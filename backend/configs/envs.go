package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT       string
	DBUSER     string
	DBHOST     string
	DBPASSWORD string
	DBPORT     string
	DBNAME     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PORT:       getEnv("PORT", "8080"),
		DBPORT:     getEnv("DBPORT", "5432"),
		DBUSER:     getEnv("DBUSER", "root"),
		DBPASSWORD: getEnv("DBPASSWORD", "mypassword"),
		DBHOST:     getEnv("DBHOST", "127.0.0.1"),
		DBNAME:     getEnv("DBNAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
