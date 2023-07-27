package token

import (
	"encoding/base64"
	"net/http"
	"strings"

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

		// Get the PASETO token from the request header
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// Extract the token from the "Bearer <token>" format
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// Verify the PASETO token with the public key
		pubKey, err := base64.RawURLEncoding.DecodeString(publicKeyStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode public key"})
			return
		}

		var accessTokenClaims paseto.JSONToken
		accessTokenFooter := map[string]interface{}{
			"type": "access",
		}

		err = paseto.NewV2().Verify(tokenStr, pubKey, &accessTokenClaims, &accessTokenFooter)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		c.Set(authorizationPayloadKey, accessTokenClaims)
		c.Next()
	}

}
