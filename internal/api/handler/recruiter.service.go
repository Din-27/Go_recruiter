package handler

import (
	"net/http"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/gin-gonic/gin"
)

func AddProfileCompany(c *gin.Context) {
	var company models.DetailPerusahaan
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
}

func GetProfileCompany(c *gin.Context) {
	var company models.DetailPerusahaan
	if err := c.ShouldBindJSON(&company); err != nil {
		_resError(c, "error", err)
		return
	}
	result := db.Create(&company)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": company})
}
