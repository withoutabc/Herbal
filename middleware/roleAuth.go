package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"herbalBody/util"
)

// CommonAuth 普通用户认证
func CommonAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok != true {
			fmt.Println("role not exist")
			util.NormErr(c, 999, "role not exist")
			c.Abort()
		}
		if role == "common" {
			c.Next()
			return
		}
		util.RespUnauthorizedErr(c)
		c.Abort()
	}
}

// AdministratorAuth 管理员认证
func AdministratorAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok != true {
			fmt.Println("role not exist")
			util.NormErr(c, 999, "role not exist")
			c.Abort()
		}
		if role == "administrator" {
			c.Next()
			return
		}
		util.RespUnauthorizedErr(c)
		c.Abort()

	}
}

// MedicAuth 医护人员认证
func MedicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if ok != true {
			fmt.Println("role not exist")
			util.NormErr(c, 999, "role not exist")
			c.Abort()
		}
		if role == "medic" {
			c.Next()
			return
		}
		util.RespUnauthorizedErr(c)
		c.Abort()
	}
}
