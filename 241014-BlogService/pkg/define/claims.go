package define

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserInfo         *UserInfo `json:"user_info"`
	RegisteredClaims *jwt.RegisteredClaims
}

// GetAudience implements jwt.Claims.
func (c UserClaims) GetAudience() (jwt.ClaimStrings, error) {
	return c.RegisteredClaims.Audience, nil
}

// GetExpirationTime implements jwt.Claims.
func (c UserClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.ExpiresAt, nil
}

// GetIssuedAt implements jwt.Claims.
func (c UserClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.IssuedAt, nil
}

// GetIssuer implements jwt.Claims.
func (c UserClaims) GetIssuer() (string, error) {
	return c.RegisteredClaims.Issuer, nil
}

// GetNotBefore implements jwt.Claims.
func (c UserClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.NotBefore, nil
}

// GetSubject implements jwt.Claims.
func (c UserClaims) GetSubject() (string, error) {
	return c.RegisteredClaims.Subject, nil
}
