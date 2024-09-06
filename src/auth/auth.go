package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GetUserIdToken returns the user id
func GetUserIdToken(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeySecret)
	if err != nil {
		return 0, err
	}

	if permissions, isOk := token.Claims.(jwt.MapClaims); isOk && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}

// GenerateToken create a token with user permissions
func GenerateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

// ValidateToken varify user has a valid token
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeySecret)
	if err != nil {
		return err
	}

	if _, isOk := token.Claims.(jwt.MapClaims); isOk && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	parts := strings.Split(token, " ")

	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}

func returnKeySecret(token *jwt.Token) (interface{}, error) {
	if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
		return nil, errors.New("invalid token")
	}

	return config.SecretKey, nil
}
