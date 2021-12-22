package utils

import (
	"os"
)

type Config struct {
	ServerPort    string
	MongoURI      string
	MongoUser     string
	MongoPassword string
	MongoDatabase string
	Debug         bool
	Testing       bool
}

func getProductionConfig() Config {
	var config Config
	config.ServerPort = getEnvWithFallback("SERVER_PORT", "8000")

	config.MongoURI = getEnvWithFallback("MONGO_URI", "mongodb://localhost:27017")
	config.MongoUser = getEnvWithFallback("MONGO_USER", "klever")
	config.MongoPassword = getEnvWithFallback("MONGO_PASSWORD", "klever")
	config.MongoDatabase = getEnvWithFallback("MONGO_DATABASE", "klever")

	config.Debug = false
	config.Testing = false
	return config
}

func getDevelopmentConfig() Config {
	config := getProductionConfig()
	config.Debug = true
	return config
}

func getTestingConfig() Config {
	config := getDevelopmentConfig()
	config.Testing = true
	return config
}

func GetConfig() Config {
	configType := GetAppEnvironment()

	switch configType {
	case "production":
		return getProductionConfig()
	case "testing":
		return getTestingConfig()
	default:
		return getDevelopmentConfig()
	}
}

func GetAppEnvironment() string {
	return getEnvWithFallback("APP_SETTINGS", "development")
}

func getEnvWithFallback(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
