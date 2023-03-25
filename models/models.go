package models

import "github.com/dgrijalva/jwt-go"



type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Name struct {
	Username string `json:"username"`
}
type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}
