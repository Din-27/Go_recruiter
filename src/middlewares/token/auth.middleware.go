package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Din-27/Go_job/helpers"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

const (
	publicKeyStr                  = "YOUR_BASE64_ENCODED_PUBLIC_KEY"
	authorizationHeaderKey        = "authorization"
	authorizationTypeBearer       = "bearer"
	authorizationPayloadKey       = "authorization_payload"
	authorizationPayloadKeyPublic = "authorization_public"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle the protected endpoint that requires a valid access token.
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
		err := paseto.NewV2().Decrypt(accessToken, helpers.SymmetricKey, &newJsonToken, &newFooter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helpers.ErrorResponse(err))
			return
		}

		c.Set(authorizationPayloadKey, newJsonToken)
		c.Next()
	}

}
