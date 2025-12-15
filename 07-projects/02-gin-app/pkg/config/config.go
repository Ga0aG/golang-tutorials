package config

import (
	"os"
)

// Config 应用配置
type Config struct {
	Port     string
	Env      string
	DBDriver string
	DBURL    string
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "sqlite"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "test.db"
	}

	return &Config{
		Port:     port,
		Env:      env,
		DBDriver: dbDriver,
		DBURL:    dbURL,
	}
}
