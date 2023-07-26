package helpers

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(id_user int, username string, email string, duration time.Duration) (string, *Payload, error)
	CreateTokenPublic(access_token string) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

var (
	c, _       = hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	PrivateKey = ed25519.PrivateKey(c)
	b, _       = hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	PublicKey  = ed25519.PublicKey(b)
)

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker() (Maker, error) {
	var symmetricKey string = "12345678901234567890123456789012"
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *PasetoMaker) CreateToken(id_user int, username string, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(id_user, username, email, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func (maker *PasetoMaker) CreateTokenPublic(access_token string) (string, error) {

	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(1 * time.Minute),
	}
	jsonToken.Set(access_token, "this is a signed message")
	footer := "some footer"
	token, err := paseto.NewV2().Sign(PrivateKey, jsonToken, footer)

	return token, err
}
