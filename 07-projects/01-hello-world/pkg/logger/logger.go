package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
)

// Init 初始化日志系统
func Init(level string) {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	warnLogger = log.New(os.Stdout, "[WARN] ", log.LstdFlags)
}

// Info 记录信息日志
func Info(msg string) {
	infoLogger.Println(msg)
}

// Error 记录错误日志
func Error(msg string) {
	errorLogger.Println(msg)
}

// Warn 记录警告日志
func Warn(msg string) {
	warnLogger.Println(msg)
}

// Infof 记录格式化信息日志
func Infof(format string, v ...interface{}) {
	infoLogger.Println(fmt.Sprintf(format, v...))
}

// Errorf 记录格式化错误日志
func Errorf(format string, v ...interface{}) {
	errorLogger.Println(fmt.Sprintf(format, v...))
}
