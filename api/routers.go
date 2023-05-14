package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/anotherone/api2"
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
	{
		ial := NewInfoApi()
		r.GET("/info/search/:user_id", ial.SearchInfo)
		r.PUT("/info/update/:user_id", ial.UpdateInfo)

		r.GET("/basic/search/:user_id", ial.SearchBasic)

		r.GET("/motor/search/:user_id", ial.SearchMotor)

		r.GET("/conclusion/search/:user_id", ial.SearchConclusion)
		r.PUT("/conclusion/update/:user_id", ial.UpdateConclusion)
	}

	//另外一边的接口
	api2.InitRouters(r)

	r.Run(":10")
}
