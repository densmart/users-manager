package jwt_auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	accessTokenTTL  = 60 * time.Minute
	refreshTokenTTL = 168 * time.Hour // 1 week
)

type JwtAuth struct {
	apiKey      string
	userID      uint64
	name        string
	email       string
	permissions map[string]uint8
}

func NewJwtAuth(apiKey string, userID uint64, userName string, email string, permissions map[string]uint8) *JwtAuth {
	return &JwtAuth{apiKey: apiKey, userID: userID, name: userName, email: email, permissions: permissions}
}

// GenerateTokens generates access & refresh tokens
func (ja *JwtAuth) GenerateTokens() (string, string, error) {
	accessToken, err := ja.GenerateAccessToken().SignedString([]byte(ja.apiKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := ja.GenerateRefreshToken().SignedString([]byte(ja.apiKey))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// GenerateAccessToken generates access token
func (ja *JwtAuth) GenerateAccessToken() *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &AccessTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ja.userID,
		ja.name,
		ja.email,
		ja.permissions,
	})
}

// GenerateRefreshToken generate refresh token
func (ja *JwtAuth) GenerateRefreshToken() *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		ja.userID,
	})
}
