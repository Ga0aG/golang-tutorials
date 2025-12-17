package infrastructure

import (
	"fmt"
	"microservice/pkg/logger"
)

// Database 数据库基础设施
type Database struct {
	connStr string
}

// NewDatabase 创建数据库连接
func NewDatabase(connStr string) *Database {
	logger.Info(fmt.Sprintf("连接数据库: %s", connStr))
	return &Database{
		connStr: connStr,
	}
}

// Connect 连接到数据库
func (db *Database) Connect() error {
	logger.Info("数据库已连接")
	return nil
}

// Close 关闭数据库连接
func (db *Database) Close() error {
	logger.Info("数据库已关闭")
	return nil
}
