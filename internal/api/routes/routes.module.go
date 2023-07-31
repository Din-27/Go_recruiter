package routes

import (
	handler "github.com/Din-27/Go_job/internal/api/handler"
	token "github.com/Din-27/Go_job/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Services(r *gin.Engine) {

	router := r.Group("/api/v1")

	router.GET("/role", handler.RoleHandle)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	router.GET("/refresh_token", handler.RefreshToken)

	authRoutes := router.Use(token.AuthMiddleware())

	authRoutes.GET("/provinsi", handler.ListProvince)
	authRoutes.GET("/kabupaten/:id_provinsi", handler.ListKabupaten)
	authRoutes.GET("/kecamatan/:id_kabupaten", handler.ListKecamatan)
	authRoutes.GET("/kelurahan/:id_kecamatan", handler.ListKelurahan)

	// USERS
	authRoutes.GET("/user", handler.GetUserById)
	authRoutes.POST("/detail/user", handler.AddUserDetail)
	authRoutes.POST("/keahlian/user", handler.AddUserKeahlian)
	authRoutes.POST("/pengalaman/user", handler.AddUserPengalaman)
	authRoutes.POST("/pendidikan/formal/user", handler.AddUserPendidikanFormal)
	authRoutes.POST("/pendidikan/nonformal/user", handler.AddUserPendidikanNonFormal)

	authRoutes.POST("/apply/lamaran/user", handler.ApplyLamaranUser)
	authRoutes.GET("/history/lamaran/user", handler.GetUserHistoryLamaranById)

	// PERUSAHAAN
	authRoutes.GET("/company", handler.GetProfileCompany)
	authRoutes.GET("/lowongan/company", handler.GetAllLowonganCompany)
	authRoutes.POST("/profile/company", handler.AddProfileCompany)

	// LOWONGAN
	authRoutes.GET("/lowongan", handler.GetAllLowongan)
	authRoutes.GET("/detail/lowongan/:id_lowongan", handler.GetDetailLowongan)
	authRoutes.POST("/lowongan", handler.AddLowongan)

	r.Run()
}
