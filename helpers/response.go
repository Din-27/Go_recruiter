package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type valueMsg struct {
}

func ResponseError(c *gin.Context, sts string, value error) {
	switch sts {
	case "server internal error":
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse(value))
	case "unauthorized":
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(value))
	case "error":
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse(value))
	}
}

func ResponseSukses(c *gin.Context, value) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"value": value})
}
