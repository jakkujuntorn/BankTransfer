package auth

import (
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

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (I_Token, error) {
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
// confrom ตาม Maker interface 
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload_Claims, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	// create token *********
	// ขั้นตอนน้อยกว่า jwt **
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	// paseto มันก็ return payload_Claims ได้เหมือนกัน
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
// confrom ตาม Maker interface 
func (maker *PasetoMaker) VerifyToken(token string) (*Payload_Claims, error) {
	payload := &Payload_Claims{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	// paseto มันก็ return payload_Claims ได้เหมือนกัน
	return payload, nil
}
