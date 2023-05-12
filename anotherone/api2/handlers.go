package api2

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouters(r *gin.Engine) {
	r.GET("")
}
func InitMyRouters() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/user/auth/openID", OpenIDLogin)
	r.Run(":" + viper.GetString("port"))
}