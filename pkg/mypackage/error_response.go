package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) *gin.H {
	return &gin.H{"message": err.Error()}
}

func ErrorReturn(value string) error {
	return errors.New(value)
}
