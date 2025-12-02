package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	// Get JWT SECRET KEY from env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT SECRET KEY is not set")
	}

	token_expiration := os.Getenv("TOKEN_EXPIRATION")
	if token_expiration == "" {
		token_expiration = "24h"
	}

	token_expiration_duration, err := time.ParseDuration(token_expiration + "h")
	if err != nil {
		return "", err
	}

	// Specify token claims
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(token_expiration_duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Create a new token with the specified signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

// Validate & parse token -> returns User ID
func ValidateToken(tokenString string) (uuid.UUID, error) {
	secret := os.Getenv("JWT_SECRET")

	// Parse token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	// Checks the type of the claims object (MapClaims)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userIDStr, ok := claims["user_id"].(string); ok {
			return uuid.Parse(userIDStr)
		}
	}

	return uuid.Nil, errors.New("invalid token claims")
}
