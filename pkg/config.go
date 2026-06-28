package pkg

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort     int
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DatabaseURL  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	readTimeOut := time.Duration(getEnvAsInt("READ_TIMEOUT", 10)) * time.Second
	writeTimeOut := time.Duration(getEnvAsInt("WRITE_TIMEOUT", 10)) * time.Second

	return &Config{
		HttpPort:     getEnvAsInt("HTTP_PORT", 8080),
		RunMode:      getEnv("RUN_MODE", "debug"),
		ReadTimeout:  readTimeOut,
		WriteTimeout: writeTimeOut,
		DatabaseURL:  getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=tekd port=5432"),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")

	if valueStr == "" {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		return fallback
	}

	return value
}
