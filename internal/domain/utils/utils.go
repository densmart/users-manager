package utils

import (
	"github.com/densmart/users-manager/internal/logger"
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

// CheckPasswordHash Check user's password in login step
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		logger.Errorf(err.Error())
		return true
	}
	return false
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
