package util

import (
	"errors"
	// "fmt"
	"time"

	 "github.com/dgrijalva/jwt-go"
	 "context"
)

const minSecretKeySize = 32
const Issuer ="ratthakorn_jakkujuntron" 

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	SecretKey string
}

// Maker is an interface for managing tokens
type I_Token interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (token string, err error)

	// AuthorizeUser แกะ token และ verify
	AuthorizeUser(ctx context.Context) (*Payload_Claims, error)

	// VerifyToken checks if the token is valid or not
	verifyToken(token string) (*Payload_Claims, error)

	
}

// NewJWTMaker creates a new JWTMaker
//  ใครจะมาใช้ให้ NewJWTMaker 
func New_JWT(SecretKey string) (I_Token, error) {
	// if len(SecretKey) < minSecretKeySize {
	// 	return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	// }
	return &JWTMaker{SecretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
// confrom ตาม Maker interface 
// func  (signature JWTMaker)CreateToken(username string, duration time.Duration) (string, *Payload_Claims, error) {
func  (signature *JWTMaker)CreateToken(username string, duration time.Duration) (token string, err error) {
	// ปั้น payload ******************************
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	// create token ***********************
	
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err = jwtToken.SignedString([]byte(signature.SecretKey))


	// SigningMethodRS512 ปลอดภัยกว่า แต่มันใช้ยังไง ******
	// key,err:=jwt.ParseRSAPrivateKeyFromPEM([]byte(signature.SecretKey))
	
	// token,_ = jwt.NewWithClaims(jwt.SigningMethodRS512, payload).SignedString([]byte(signature.SecretKey))
	
	
	return token, err
}

// VerifyToken checks if the token is valid or not
// confrom ตาม Maker interface 
func  (signature *JWTMaker)verifyToken(token string) (*Payload_Claims, error) {
	
	
	// สร้าง  key func สำหรับ return signature******
	//ParseWithClaims กับ Parse ต้องใช้มัน *********
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(signature.SecretKey), nil
	}

	// ParseWithClaims รับ paramiter 3 ตัว (token,Claims,keyFunc)
	//Parse รับ paramiter 2 ตัว  (token,keyFunc)
	jwtToken, err := jwt.ParseWithClaims(token, &Payload_Claims{}, keyFunc)
	if err != nil  {
		// เช็ค error ของ JWT เหรอ  ทำไมเช็ค เยอะจัง********
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	//ParseWithClaims  .Claims(ใส Payload_Claims ที่เราสร้างขึ้น)
	//Parse  .Claims(ใส jwt.MapClaims เป็นของ jwt) ลองใส *Payload_Claims ก็ได้ไม่ error ****

	payload, ok := jwtToken.Claims.(*Payload_Claims)

	// แบบนี้ error  เพราะ func นี้ให้ retur *Payload_Claims ***
	// payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

	

