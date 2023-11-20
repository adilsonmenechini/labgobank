package utils

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(id, email string) (string, error) {

	expirationTime := time.Now().Add(30 * time.Second)
	claims := &Claims{
		ID:    id,
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

func SetTokenAsCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

func GetTokenFromCookie(r *http.Request) (Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return Claims{}, err
	}

	ctk, err := ValidateToken(cookie.Value)
	if err != nil {
		return Claims{}, err
	}
	return Claims{
		ID:    ctk.ID,
		Email: ctk.Email,
		Name:  ctk.Name,
	}, nil
}
