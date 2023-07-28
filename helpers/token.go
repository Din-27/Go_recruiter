package helpers

import (
	"crypto/ed25519"
	"encoding/hex"
	"os"
	"time"

	"github.com/o1egl/paseto"
)

var SymmetricKey = []byte(os.Getenv("SYMETRIC_KEY")) // Must be 32 bytes

const (
	tokenDuration = 15 * time.Minute // Adjust as per your requirements
)

func GenerateAccessToken(username, email, role string, durasi time.Duration) (string, error) {
	now := time.Now()
	exp := now.Add(durasi)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   username,
		Issuer:     email,
		Subject:    role,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim    to the token
	jsonToken.Set("data", "this is a signed message")
	footer := "access"

	// Encrypt data
	token, err := paseto.NewV2().Encrypt(SymmetricKey, jsonToken, footer)
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
