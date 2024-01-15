package auth

import (
	"time"
)

// Maker is an interface for managing tokens
type I_Token interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload_Claims, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload_Claims, error)
}
