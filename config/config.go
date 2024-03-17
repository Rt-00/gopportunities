package config

import (
	"gorm.io/gorm"
)

// Var fora de uma função é um variavel de todo o package
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
