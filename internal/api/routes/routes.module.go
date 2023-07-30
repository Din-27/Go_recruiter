package routes

import (
	handler "github.com/Din-27/Go_job/internal/api/handler"
	token "github.com/Din-27/Go_job/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Services(r *gin.Engine) {

	router := r.Group("/api/v1")

	router.GET("/role", handler.RoleHandle)

	router.POST("/register/:role", handler.Register)
	router.POST("/login/:role", handler.Login)

	router.GET("/refresh_token", handler.RefreshToken)

	authRoutes := router.Use(token.AuthMiddleware())

	authRoutes.GET("/provinsi", handler.ListProvince)
	authRoutes.GET("/kabupaten/:id_provinsi", handler.ListKabupaten)
	authRoutes.GET("/kecamatan/:id_kabupaten", handler.ListKecamatan)
	authRoutes.GET("/kelurahan/:id_kecamatan", handler.ListKelurahan)

	authRoutes.GET("/user", handler.GetUserById)
	authRoutes.POST("/detail/user", handler.AddUserDetail)
	authRoutes.GET("/apply", handler.ApplyLamaranUser)
	r.Run()
}
