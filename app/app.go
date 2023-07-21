package app

import (
	"log"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/controllers"
)

func AppRoutes() {
	tokenMaker, err := helpers.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatal(err)
	}
	controllers.Services(tokenMaker)
}
