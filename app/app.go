package app

import (
	"log"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes() {
	tokenMaker, err := helpers.NewPasetoMaker()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	controllers.Services(tokenMaker, router)
}
