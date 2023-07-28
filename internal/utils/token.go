package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"os"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/argon2"
)

var (
	value        = os.Getenv("SYMETRIC_KEY")
	symmetricKey = []byte(value) // Must be 32 bytes
	Key          = argon2.IDKey([]byte(symmetricKey), []byte("asdoiwje#"), 1, 64*1024, 4, 32)
)

const (
	tokenDuration = 15 * time.Minute // Adjust as per your requirements
)

func GenerateAccessToken(username, email, role string, durasi time.Duration) (string, error) {
	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(durasi),
	}
	jsonToken.Set("username", username)
	jsonToken.Set("email", email)
	jsonToken.Set("role", role)
	footer := "access"

	// Encrypt data
	token, err := paseto.NewV2().Encrypt(Key, jsonToken, footer)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(username, email, role string, durasi time.Duration) (string, error) {
	b, _ := hex.DecodeString(os.Getenv("PRIVATE_KEY"))
	privateKey := ed25519.PrivateKey(b)

	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(durasi),
	}

	// Add custom claim    to the token
	jsonToken.Set("username", username)
	jsonToken.Set("email", email)
	jsonToken.Set("role", role)
	footer := "refresh"

	// Sign data
	token, err := paseto.NewV2().Sign(privateKey, jsonToken, footer)

	return token, err
}
