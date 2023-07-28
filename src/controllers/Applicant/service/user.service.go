package service

import (
	"net/http"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/config"
	"github.com/gin-gonic/gin"
)

var (
	db        = config.DBinit()
	_resError = helpers.ResponseError
	_isErr    = helpers.ErrorReturn
)

func User(c *gin.Context) {
	value, err := helpers.FetchGetProvinsi()
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}
