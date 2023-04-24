package utils

import (
	"github.com/spf13/viper"
	"net/url"
	"rsc.io/qr"
)

// Create2faURL Create URL for 2FA application
func Create2faURL(email, secret string) (string, error) {
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

// CreateQRCode Creates QR code image for 2FA
func CreateQRCode(data2code string) ([]byte, error) {
	code, err := qr.Encode(data2code, qr.Q)
	if err != nil {
		return nil, err
	}
	data := code.PNG()
	return data, err
}
