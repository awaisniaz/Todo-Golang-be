package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenExpireDuration  = time.Minute * 15
	refreshTokenExpireDuration = time.Hour * 24 * 7 // 7 days
	secretKey                  = "todolist"
)

// Custom claims for tokens
type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(Id string) (string, error) {
	claims := CustomClaims{
		UserID: Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to parse token: %w", err)
	}
	fmt.Println(token)
	if token.Valid == true {
		return true, nil

	}

	return false, nil
}
func HashPassword(pass string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
