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
		u.POST("/password", middleware.JWTAuthMiddleware(), uai.ChangePassword)
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
		r.GET("/info/search/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.SearchInfo)
		r.PUT("/info/update/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.UpdateInfo)

		r.POST("/basic/add/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.AddBasic)
		r.GET("/basic/search/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.SearchBasic)

		r.POST("/motor/add/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.AddMotor)
		r.GET("/motor/search/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.SearchMotor)

		r.GET("/conclusion/search/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.SearchConclusion)
		r.PUT("/conclusion/update/:user_id", middleware.JWTAuthMiddleware(), middleware.CommonAuth(), ial.UpdateConclusion)
	}

	//另外一边的接口
	api2.InitRouters(r)

	r.Run(":10")
}
