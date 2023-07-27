package app

import (

	"github.com/Din-27/Go_job/src/routes"
	"github.com/gin-gonic/gin"
)

func AppRoutes() {
	
	router := gin.Default()
	routes.Services(router)
}
