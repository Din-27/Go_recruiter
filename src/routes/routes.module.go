package routes

import (
	"os"

	users "github.com/Din-27/Go_job/src/controllers/applicant/service"
	auth "github.com/Din-27/Go_job/src/controllers/auth/service"

	// company "github.com/Din-27/Go_job/src/controllers/recruiter/service"
	states "github.com/Din-27/Go_job/src/controllers/state/service"

	"github.com/Din-27/Go_job/src/middlewares/token"
	"github.com/gin-gonic/gin"
)

func Services(r *gin.Engine) {

	router := r.Group("/api/v1")

	router.GET("/role", auth.RoleHandle)

	router.POST("/register/:role", auth.Register)
	router.POST("/login/:role", auth.Login)

	router.POST("/refresh_token", auth.RefreshToken)

	authRoutes := router.Use(token.AuthMiddleware())

	authRoutes.GET("/provinsi", states.ListProvince)
	authRoutes.GET("/kabupaten/:id_provinsi", states.ListKabupaten)
	authRoutes.GET("/kecamatan/:id_kabupaten", states.ListKecamatan)
	authRoutes.GET("/kelurahan/:id_kecamatan", states.ListKelurahan)

	authRoutes.GET("/user", users.User)

	r.Run(os.Getenv("PORT"))
}
