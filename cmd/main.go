package main

import (
	"github.com/Din-27/Go_job/internal/api/routes"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()
	routes.Services(router)
}
