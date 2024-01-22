package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload_Claims struct {
	jwt.StandardClaims
}

// StandardClaims **************
// หาข้อมูลว่าควรใสอะไรลงปบ้าง *****************
// type StandardClaims struct {
// 	Audience  string `json:"aud,omitempty"` // User
// 	ExpiresAt int64  `json:"exp,omitempty"` // ระยะเวลา
// 	Id        string `json:"jti,omitempty"` // Id
// 	IssuedAt  int64  `json:"iat,omitempty"` //  ออกที่ไหน
// 	Issuer    string `json:"iss,omitempty"`// ใครเป็นคนออก
// 	NotBefore int64  `json:"nbf,omitempty"`
// 	Subject   string `json:"sub,omitempty"`
// }

// NewPayload creates a new token payload with a specific username and duration
// ทำหน้าที่ปั้น payload ******
func NewPayload(username string, duration time.Duration) (*Payload_Claims, error) {
	// ปั้น payload ********
	// ปั้นแบบนี้ไม่ได้เพรา อาจจะมาจาก lib รึื ป่าว
	// payload := &Payload_Claims{
	// 	Audience:  "",
	// 	ExpiresAt: int64(4),
	// }

	// ปั้น payload ********
	payload := &Payload_Claims{}

	payload.Audience = username
	payload.ExpiresAt = time.Now().Add(duration).Unix()
	payload.IssuedAt = time.Now().Unix()
	payload.Issuer = Issuer

	return payload, nil
}

// Valid checks if the token payload is valid or not
// ทำให้เป็น jwt.Cliams เท่านั้น
//  Func นี้ทำงานอะไรบ้าง
// func (payload *Payload_Claims) Valid() error {
// 	if time.Now().After(payload.ExpiresAt) {
// 		return ErrExpiredToken
// 	}
// 	return nil
// }
