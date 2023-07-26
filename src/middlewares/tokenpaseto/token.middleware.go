package tokenpaseto

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Din-27/Go_job/helpers"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

const (
	authorizationHeaderKey        = "authorization"
	authorizationTypeBearer       = "bearer"
	authorizationPayloadKey       = "authorization_payload"
	authorizationPayloadKeyPublic = "authorization_public"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddlewareLocal(tokenMaker helpers.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}

func AuthMiddlewarePublic(tokenMaker helpers.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newJsonToken paseto.JSONToken
		var newFooter string
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		err := paseto.NewV2().Verify(accessToken, helpers.PublicKey, &newJsonToken, &newFooter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}
		if time.Now().After(newJsonToken.Expiration) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token has expired"})
			return
		}
		c.Set(authorizationPayloadKeyPublic, newJsonToken)
		c.Next()
	}
}
