package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type RegisterUser struct {
	UserId          int    `json:"user_id" form:"user_id"`
	Username        string `json:"username" form:"username" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

type LoginUser struct {
	UserId   int    `json:"user_id" form:"user_id"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}

type MyClaims struct {
	UserId    int
	LoginTime time.Time
	Type      string
	jwt.StandardClaims
}
