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
	Role            string `json:"role" form:"role" binding:"required"` //角色，详见roleAuth.go
}

type LoginUser struct {
	UserId   int    `json:"user_id" form:"user_id"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}

type ReqChangePwd struct {
	Password string `json:"password" form:"password" binding:"required"`
}

type MyClaims struct {
	UserId    int
	Role      string    //角色，详见roleAuth.go
	LoginTime time.Time //登录时间
	Type      string    //token类型
	jwt.StandardClaims
}
