package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"os"

	"github.com/Din-27/Go_job/internal/models"
	"github.com/o1egl/paseto"
)

func DecodedToken(token string) (data models.Decoded, err error) {
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
