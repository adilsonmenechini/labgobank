package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

type Claims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id, name, email string) (string, error) {

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (Claims, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return Claims{}, err
		}
		return Claims{}, err
	}
	if !tkn.Valid {
		return Claims{}, err
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return Claims{}, errors.New("Token Expired")
	}

	return Claims{
		ID:    claims.ID,
		Email: claims.Email,
		Name:  claims.Name,
	}, nil
}

func SetTokenAuthorization(w http.ResponseWriter, token string) {
	w.Header().Set("Authorization", "Bearer "+token)
}

func GetTokenAuthorization(r *http.Request) (Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return Claims{}, errors.New("Authorization header not found")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	ctk, err := ValidateToken(tokenString)
	if err != nil {
		return Claims{}, err
	}

	return Claims{
		ID:    ctk.ID,
		Email: ctk.Email,
		Name:  ctk.Name,
	}, nil
}
