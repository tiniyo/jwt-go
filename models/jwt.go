package models

import (
	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtClaims struct {
	JwtInfo
	jwt.StandardClaims
}

type JwtInfo struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}
