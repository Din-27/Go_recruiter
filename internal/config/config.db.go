package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBinit() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}