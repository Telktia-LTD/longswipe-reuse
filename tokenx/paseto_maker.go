package tokenx

import (
	"fmt"
	"time"

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"github.com/sirupsen/logrus"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(user interfacesx.UserResponse, duration time.Duration) (string, *AuthPayload, error) {
	payload, err := NewPayload(user, duration)
	if err != nil {
		return "", payload, err
	}
	logrus.Info("payload", payload)

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*AuthPayload, error) {
	payload := &AuthPayload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
