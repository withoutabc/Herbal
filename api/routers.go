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
		uai := NewUserApi()
		u.POST("/register", uai.Register)
		u.POST("/login", uai.Login)
		u.POST("/refresh", middleware.JWTAuthMiddleware(), uai.Refresh)
	}

	r.GET("/questionnaire", GetQuestionnaire)
	{
		qal := NewSubmissionApi()
		r.POST("/submit", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), qal.ReceiveSubmission)
		eal := NewExcelApi()
		r.POST("/gen/excel/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), eal.GenExcel)
		r.POST("/upload/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), Upload)
		r.GET("/comment/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), qal.GetComment)
	}
	r.GET("/excels/zips", GetExcelZips)

	{
		pal := NewPromiseApi()
		r.POST("/promise", middleware.JWTAuthMiddleware(), pal.QueryPromises)
	}
	r.Run(":10")
}
