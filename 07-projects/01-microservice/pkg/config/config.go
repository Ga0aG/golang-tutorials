package config

import (
	"os"
	"strconv"
)

// Config 应用配置
type Config struct {
	Port     int
	LogLevel string
	Database string
	Env      string
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	port := 8080
	if p := os.Getenv("PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	database := os.Getenv("DATABASE_URL")
	if database == "" {
		database = "localhost:5432"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		Port:     port,
		LogLevel: logLevel,
		Database: database,
		Env:      env,
	}
}
