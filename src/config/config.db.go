package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBinit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/job_website"), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}