package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload_Claims struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
// ทำหน้าที่ปั้น payload ******
func NewPayload(username string, duration time.Duration) (*Payload_Claims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// ปั้น payload
	payload := &Payload_Claims{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
// ทำให้เป็น jwt.Cliams เท่านั้น
//  Func นี้ทำงานอะไรบ้าง
func (payload *Payload_Claims) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
