package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
