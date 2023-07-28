package main

import (
	"fmt"
	"log"

	"github.com/Din-27/Go_job/internal/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file %s", err))
	}
	router := gin.Default()
	routes.Services(router)
}
