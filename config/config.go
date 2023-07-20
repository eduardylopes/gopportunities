package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	l  *Logger
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	l = NewLogger(p)
	return l
}

func GetSQLite() *gorm.DB {
	return db
}
