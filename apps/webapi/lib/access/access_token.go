package access

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"provar.se/webapi/lib/user"
)

const (
	accessTokenTTL = 24 * time.Hour
)

var (
	hmacSecret = []byte("your-secret-key")

	ErrInvalidToken = errors.New("invalid token")
)

// Setup sets the secret key used to sign the JWT tokens
func Setup(secret string) error {
	hmacSecret = []byte(secret)
	return nil
}

// CreateAccessToken creates a new access token for the given user id
func CreateAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(accessTokenTTL).Unix(),
	})
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateAccessToken validates the given access token and returns the user id
func ValidateAccessToken(tokenString string) (*user.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	userId, ok := claims["userId"].(string)
	if !ok {
		return nil, ErrInvalidToken
	}
	return user.FindByID(userId)
}
