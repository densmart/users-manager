package utils

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"github.com/dgryski/dgoogauth"
	"github.com/spf13/viper"
	"net/url"
)

func CreateOtpSecret() (string, error) {
	secret := make([]byte, 10)
	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(secret), nil
}

func CreateOtpURL(email, secret string) (string, error) {
	issuer := viper.GetString("app.name")
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		return "", err
	}

	URL.Path += "/" + url.PathEscape(email)

	params := url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", issuer)

	URL.RawQuery = params.Encode()

	return URL.String(), err
}

func VerifyOTP(token, secret string) error {
	otpConfig := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  5,
		HotpCounter: 0,
	}

	val, err := otpConfig.Authenticate(token)
	if err != nil {
		return err
	}

	if !val {
		return errors.New("OTP key is invalid")
	}

	return nil
}
