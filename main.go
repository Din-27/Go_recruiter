package main

import (
	"log"

	users "github.com/Din-27/Go_job/controllers/auth/service"
	"github.com/Din-27/Go_job/middlewares/tokenpaseto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/job_website"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error !")
	}

	tokenMaker, err := tokenpaseto.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	userRepository := users.NewRepositoryUser(db)
	router := r.Group("/api/v1")
	router.POST("/register", userRepository.Register)
	router.POST("/login", userRepository.Login)
	authRoutes := router.Group("/").Use(tokenpaseto.AuthMiddleware(tokenMaker))
	authRoutes.GET("/test", userRepository.Test)

	r.Run()
}
