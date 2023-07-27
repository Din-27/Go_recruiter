package service

import (
	"net/http"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/Recruiter/schema"
	"github.com/gin-gonic/gin"
)

var (
	db        = config.DBinit()
	_resError = helpers.ResponseError
	_isErr    = helpers.ErrorReturn
)

func AddProfileCompany(c *gin.Context) {
	var company schema.DetailPerusahaan
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
	var company schema.DetailPerusahaan
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
