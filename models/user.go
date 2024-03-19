package models

import (
	"github.com/golang-jwt/jwt"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email       string `json:"email"`
	AccountType string `json:"accountType"`
	jwt.StandardClaims
}

type Session struct {
	Email       string `json:"email"`
	Token       string `json:"token"`
	AccountType string `json:"accountType"`
}

type User struct {
	ID          string  `json:"id" bson:"id" validate:"required" auto:"uuid" unique:"true"`
	Email       string  `json:"email" bson:"email" validate:"required,email" unique:"true"`
	Password    string  `json:"-" bson:"password" validate:"required" omitempty:"true"`
	AccountType string  `json:"accountType" bson:"accountType" validate:"required,oneof=admin user"`
	FirstName   *string `json:"firstName" bson:"firstName" omitempty:"true"`
	LastName    *string `json:"lastName" bson:"lastName" omitempty:"true"`
	Phone       *string `json:"phone" bson:"phone" omitempty:"true"`
	Address     *string `json:"address" bson:"address" omitempty:"true"`
	City        *string `json:"city" bson:"city" omitempty:"true"`
	Country     *string `json:"country" bson:"country" omitempty:"true"`
	PostalCode  *string `json:"postalCode" bson:"postalCode" omitempty:"true"`
}
