package config

import (
	"os"

	// "github.com/Din-27/Go_job/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBinit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	db, err := gorm.Open(mysql.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		return nil
	}
	// db.AutoMigrate(&models.Perusahaan{})
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.DetailPerusahaan{})
	// db.AutoMigrate(&models.DetailUser{})
	// db.AutoMigrate(&models.PendidikanFormalUser{})
	// db.AutoMigrate(&models.PendidikanNonFormalUser{})
	// db.AutoMigrate(&models.PengalamanUser{})
	// db.AutoMigrate(&models.DetailPerusahaan{})
	return db
}
