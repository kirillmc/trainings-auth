package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
	Role  Role   `json:"role"`
}
