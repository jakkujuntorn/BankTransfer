package auth

import (
	"errors"
	"fmt"
	"time"

	 "github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
//  เช็คจำนวน secretKey
// มันทำหน้าที่อะไร *****
func NewJWTMaker(secretKey string) (I_Token, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
// confrom ตาม Maker interface 
func  CreateToken(username string, duration time.Duration) (string, *Payload_Claims, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	// create token ********
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
// confrom ตาม Maker interface 
func  VerifyToken(token string) (*Payload_Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload_Claims{}, keyFunc)
	if err != nil  {
		// เช็ค error ของ JWT เหรอ  ทำไมเช็ค เยอะจัง********
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload_Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
