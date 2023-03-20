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
		u.POST("/refresh", Refresh, middleware.JWTAuthMiddleware())
	}

	r.POST("/submit", ReceiveSubmission, middleware.JWTAuthMiddleware())
	r.GET("/questionnaire", GetQuestionnaire)
	r.Run(":80")
}
