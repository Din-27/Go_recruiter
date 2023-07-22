package controllers

import (
	"github.com/Din-27/Go_job/helpers"
	users "github.com/Din-27/Go_job/src/controllers/auth/service"
	"github.com/Din-27/Go_job/src/middlewares/tokenpaseto"
	"github.com/gin-gonic/gin"
)

func Services(tokenMaker helpers.Maker, r *gin.Engine) {

	router := r.Group("/api/v1")

	router.POST("/register", users.Register)
	router.POST("/login", users.Login)
	authRoutes := router.Group("/").Use(tokenpaseto.AuthMiddleware(tokenMaker))
	authRoutes.GET("/checkauth", users.RefreshToken)

	r.Run()
}
