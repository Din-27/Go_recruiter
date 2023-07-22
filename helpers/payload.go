package helpers

import (
	"errors"
	"time"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        int       `json:"id_user"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ApiToken  string    `json:"apitoken"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(id_user int, username string, email string, apitoken string, duration time.Duration) (*Payload, error) {

	payload := &Payload{
		ID:        id_user,
		Username:  username,
		Email:     email,
		ApiToken:  apitoken,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func Valid(payload *Payload) error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
