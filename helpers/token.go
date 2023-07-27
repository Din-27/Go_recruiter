package helpers

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

const (
	PublicKeyStr  = "YOUR_BASE64_ENCODED_PUBLIC_KEY"
	PrivateKeyStr = "YOUR_BASE64_ENCODED_PRIVATE_KEY"
	tokenDuration = 15 * time.Minute // Adjust as per your requirements
)

func PrivateKeyHandle() []byte {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating key pair:", err)
		return nil
	}
	privKey := base64.RawURLEncoding.EncodeToString(privateKey)
	if err != nil {
		fmt.Println("Error decoding private key:", err)
		return nil
	}
	privateKey = ed25519.PrivateKey(privKey)
	return privateKey
}

func GenerateAccessToken(username, email, role string) (string, error) {
	privateKey := PrivateKeyHandle()
	now := time.Now()
	expiration := now.Add(15 * time.Minute)

	// Create a new PASETO token for the access token
	accessToken := paseto.JSONToken{
		Subject:    username,
		Audience:   email,
		Issuer:     role,
		Expiration: expiration,
	}
	accessTokenFooter := map[string]interface{}{
		"type": "access",
	}

	// Sign the access token with the private key

	// Sign the access token with the private key
	token, err := paseto.NewV2().Sign(privateKey, &accessToken, accessTokenFooter)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken() (string, error) {
	privateKey := PrivateKeyHandle()
	refreshTokenBytes := make([]byte, 32)
	_, err := rand.Read(refreshTokenBytes)
	if err != nil {
		return "", err
	}

	// Create a new PASETO token for the refresh token
	refreshToken := paseto.JSONToken{}
	refreshTokenFooter := map[string]interface{}{"type": "refresh"}

	// Sign the refresh token with the private key
	token, err := paseto.NewV2().Sign(privateKey, &refreshToken, refreshTokenFooter)
	if err != nil {
		return "", err
	}

	return token, nil
}
