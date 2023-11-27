package config

import (
	"fmt"
	"os"

	// "github.com/Din-27/Go_job/internal/models"
	"github.com/Din-27/Go_job/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBinit() *gorm.DB {
	utils.LoadEnv()
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@/%v", "root", "", "job_website")), &gorm.Config{})
		if err != nil {
			return nil
		}
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
