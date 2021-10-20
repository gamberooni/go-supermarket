package model

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID          int
	Name        string `json:"name" validate:"required"`
	Birthday    string `json:"birthday" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
