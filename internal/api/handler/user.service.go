package handler

import (
	"fmt"
	"net/http"

	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	_isvalue, _ := c.Get("authorization_payload")
	fmt.Println(_isvalue)
	value, err := utils.FetchGetProvinsi()
	if err != nil {
		_resError(c, "error", err)
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}
