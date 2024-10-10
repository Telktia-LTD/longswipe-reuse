package tokenx

import (
	"errors"
	"time"

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type AuthPayload struct {
	User      interfacesx.UserResponse `json:"user"`
	IssuedAt  time.Time                `json:"iat"`
	ExpiresAt time.Time                `json:"exp"`
}

func NewPayload(user interfacesx.UserResponse, duration time.Duration) (*AuthPayload, error) {
	payload := &AuthPayload{
		User:      user,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *AuthPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
