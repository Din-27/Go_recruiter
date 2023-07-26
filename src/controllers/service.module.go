package controllers

import (
	"github.com/Din-27/Go_job/helpers"
	auth "github.com/Din-27/Go_job/src/controllers/auth/service"
	keahlian "github.com/Din-27/Go_job/src/controllers/keahlian/service"
	states "github.com/Din-27/Go_job/src/controllers/state"
	users "github.com/Din-27/Go_job/src/controllers/user/service"

	"github.com/Din-27/Go_job/src/middlewares/tokenpaseto"
	"github.com/gin-gonic/gin"
)

func Services(tokenMaker helpers.Maker, r *gin.Engine) {

	router := r.Group("/api/v1")

	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	authRoutes := router.Group("/").Use(tokenpaseto.AuthMiddlewareLocal(tokenMaker))

	authRoutes.GET("/refresh_token", auth.RefreshToken)

	authRoutesPublic := router.Group("/").Use(tokenpaseto.AuthMiddlewarePublic(tokenMaker))

	authRoutesPublic.GET("/provinsi", states.ListProvince)
	authRoutesPublic.GET("/kabupaten/:id_provinsi", states.ListKabupaten)
	authRoutesPublic.GET("/kecamatan/:id_kabupaten", states.ListKecamatan)
	authRoutesPublic.GET("/kelurahan/:id_kecamatan", states.ListKelurahan)

	authRoutesPublic.GET("/keahlian", keahlian.ListKeahlian)
	authRoutesPublic.POST("/keahlian", keahlian.AddKeahlian)

	authRoutesPublic.GET("/user", users.User)

	r.Run()
}
