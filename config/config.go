package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitializeConfigurations() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return logger.Errorf("Config initialization error: ", err)
	}

	return nil
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)

	return logger
}
