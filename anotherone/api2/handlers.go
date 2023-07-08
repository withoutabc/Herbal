package api2

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.GET("/user/auth/openID", OpenIDLogin)
	r.GET("/user/checkphone", CheckPhoneTest)
	r.GET("/appidandappsecret", getAppIdAndAppSecret)
}
