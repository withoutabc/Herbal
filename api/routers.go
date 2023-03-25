package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.POST("/refresh", middleware.JWTAuthMiddleware(), Refresh)
	}
	r.GET("/questionnaire", GetQuestionnaire)
	{
		r.POST("/submit", middleware.JWTAuthMiddleware(), ReceiveSubmission)
		r.GET("/excel/:user_id", middleware.JWTAuthMiddleware(), GetExcel)
		r.POST("/upload/:user_id", middleware.JWTAuthMiddleware(), Upload)
		r.GET("/comment/:user_id", middleware.JWTAuthMiddleware(), Comment)
	}
	r.Run(":10")
}
