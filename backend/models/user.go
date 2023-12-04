package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// Post Response JSON
type PostUser struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	AdminEmail string `json:"admin_email"`
}

type UserId struct {
	AdminEmail string `json:"admin_email"`
	Id         int    `json:"id"`
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type UserJWT struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
