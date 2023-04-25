package utils

import (
	"crypto/rand"
	"encoding/base32"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
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

// BinaryStringToDecimal Convert binary string like "10111" to decimal number
func BinaryStringToDecimal(binString string) (int, error) {
	var remainder int
	binCode, err := strconv.Atoi(binString)
	if err != nil {
		return 0, err
	}
	index := 0
	decimalNum := 0
	for binCode != 0 {
		remainder = binCode % 10
		binCode = binCode / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum, nil
}
