package util

import (
	"fmt"
	"net/mail"
	"regexp"
	"errors"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func Validate_MoneyValue(moneyValue int32) error {

	if moneyValue <= 0 {
		return errors.New("money is negative or zero")
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	// ถ้าไม่เท่ากับ แสดงว่า มี error ไม่แมชกับ regexp
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	// ถ้าไม่เท่ากับ แสดงว่า มี error ไม่แมชกับ regexp
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	// return อะไรออกมา และเอาไว้ทำอะไร
	// เช็ค pattern email ***
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}


// func ValidateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

// 	// pb.CreateUserRequest จะดึงค่าตัวแปรที่สร้างไว้กับ proto ได้  username, password, email, fullname 
// 	if err := ValidateUsername(req.GetUsername()); err != nil {
// 		// ส่งข้อมูลไปปั้น error ใหม่ **
// 		violations = append(violations, fieldViolation("username", err))
// 	}

// 	// Password
// 	if err := ValidatePassword(req.GetPassword()); err != nil {
// 		violations = append(violations, fieldViolation("password", err))
// 	}

// 	// Fullanem
// 	if err := ValidateFullName(req.GetFullName()); err != nil {
// 		violations = append(violations, fieldViolation("full_name", err))
// 	}

// 	// Email
// 	if err := ValidateEmail(req.GetEmail()); err != nil {
// 		violations = append(violations, fieldViolation("email", err))
// 	}

// 	return violations
// }