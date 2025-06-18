package config

import (
	"os"
	_ "os"

	_ "github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "8080"),
		AppURL:       getEnv("APP_URL", "http://localhost:8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5430"),
		DBUser:       getEnv("DB_USER", "cmulifelongedtemp"),
		DBPassword:   getEnv("DB_PASS", "12345"),
		DBName:       getEnv("DB_NAME", "cmu_lifelong_ed_db_temp"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", "helloworld"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
