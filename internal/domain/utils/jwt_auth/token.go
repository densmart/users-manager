package jwt_auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	apiKey  string
	Access  string
	Refresh string
}

func NewJwtToken(apiKey string) *JwtToken {
	return &JwtToken{apiKey: apiKey}
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID      uint64 `json:"user_id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Permissions map[string]uint8
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID uint64 `json:"user_id"`
}

// GetAccessClaims checks access authentication and returns access token claims
func (tc *JwtToken) GetAccessClaims() (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tc.Access, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tc.apiKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid access token claims")
	}
	return claims, nil
}

// GetRefreshClaims checks refresh authentication and returns refresh token claims
func (tc *JwtToken) GetRefreshClaims() (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tc.Refresh, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(tc.apiKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}
	return claims, nil
}
