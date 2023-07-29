package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
)

const (
	authorizationHeaderKey        = "Authorization"
	authorizationTypeBearer       = "bearer"
	authorizationPayloadKey       = "authorization_payload"
	authorizationPayloadKeyPublic = "authorization_public"
)

func RefreshDecodedToken(token string) (data models.Decoded, err error) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	b, _ := hex.DecodeString(os.Getenv("PUBLIC_KEY"))
	publicKey := ed25519.PublicKey(b)
	if err := paseto.NewV2().Verify(token, publicKey, &newJsonToken, &newFooter); err != nil {
		return data, err
	}
	data.Username = newJsonToken.Get("username")
	data.Email = newJsonToken.Get("email")
	data.Role = newJsonToken.Get("role")
	return data, nil
}

func DecodedTokenBearer(c *gin.Context) (data models.Decoded, err error) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	authorizationHeader := c.GetHeader(authorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		return data, err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return data, err
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		return data, err
	}

	accessToken := fields[1]
	err = paseto.NewV2().Decrypt(accessToken, Key, &newJsonToken, &newFooter)
	if err != nil {
		return data, err
	}
	data.Username = newJsonToken.Get("username")
	data.Email = newJsonToken.Get("email")
	data.Role = newJsonToken.Get("role")
	return data, nil
}
