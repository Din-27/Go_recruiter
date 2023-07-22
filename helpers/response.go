package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, sts string, value error) {
	switch sts {
	case "server internal error":
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse(value))
	case "unauthorized":
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(value))
	case "error":
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse(value))
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse(value))
	}
}
