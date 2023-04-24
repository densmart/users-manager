package utils

import (
	"crypto/rand"
	"encoding/base32"
	"golang.org/x/crypto/bcrypt"
)

func Divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

// GeneratePasswordHash Create hash from user's given password to save in database
// Using bcrypt package to create hash at the default cost
func GeneratePasswordHash(password string) (string, error) {
	pwd := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Create2FASecret Create random string secret for 2 factor authenticator app
func Create2FASecret() (string, error) {
	secret := make([]byte, 10)
	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(secret), nil
}
