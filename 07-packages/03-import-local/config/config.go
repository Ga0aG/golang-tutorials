package config

import "fmt"

// Config 配置结构体
type Config struct {
    Host     string
    Port     int
    Database string
    Username string
    Password string
}

// LoadConfig 加载配置
func LoadConfig() Config {
    return Config{
        Host:     "localhost",
        Port:     5432,
        Database: "mydb",
        Username: "admin",
        Password: "secret",
    }
}

// GetDatabaseURL 获取数据库连接URL
func (c Config) GetDatabaseURL() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
        c.Username, c.Password, c.Host, c.Port, c.Database)
}