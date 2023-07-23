package keahlian

import (
	"net/http"

	"github.com/Din-27/Go_job/helpers"
	"github.com/Din-27/Go_job/src/config"
	"github.com/Din-27/Go_job/src/controllers/keahlian/schema"
	"github.com/gin-gonic/gin"
)

var (
	db        = config.DBinit()
	_resError = helpers.ResponseError
	_isErr    = helpers.ErrorReturn
)

func ListKeahlian(c *gin.Context) {
	var keahlian []schema.Keahlian
	result := db.Find(&keahlian)
	if result.Error != nil {
		_resError(c, "error", _isErr("Email tidak ditemukan !"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": keahlian})
}

func AddKeahlian(c *gin.Context) {
	var keahlian schema.Keahlian
	if err := c.ShouldBindJSON(&keahlian); err != nil {
		_resError(c, "error", err)
		return
	}
	result := db.Create(&keahlian)
	if result.Error != nil {
		_resError(c, "server internal error", result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": keahlian})
}
