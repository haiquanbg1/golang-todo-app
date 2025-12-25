package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err := checkDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func checkDB(db *gorm.DB) error {
	dbInstance, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := dbInstance.Ping(); err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	return nil
}
