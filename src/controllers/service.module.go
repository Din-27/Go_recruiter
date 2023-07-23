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
	
	authRoutes := router.Group("/").Use(tokenpaseto.AuthMiddleware(tokenMaker))
	
	authRoutes.GET("/checkauth", auth.RefreshToken)

	authRoutes.GET("/provinsi", states.ListProvince)
	authRoutes.GET("/kabupaten/:id_provinsi", states.ListKabupaten)
	authRoutes.GET("/kecamatan/:id_kabupaten", states.ListKecamatan)
	authRoutes.GET("/kelurahan/:id_kecamatan", states.ListKelurahan)

	authRoutes.GET("/keahlian", keahlian.ListKeahlian)
	authRoutes.POST("/keahlian", keahlian.AddKeahlian)

	authRoutes.GET("/user", users.User)

	r.Run()
}
