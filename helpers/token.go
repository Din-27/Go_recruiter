package helpers

import (
	"crypto/ed25519"
	"encoding/hex"
	"time"

	"github.com/o1egl/paseto"
)

var SymmetricKey = []byte("YELLOW SUBMARINE, BLACK WIZARDRY") // Must be 32 bytes

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
	b, _ := hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
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
