package api2

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouters(r *gin.Engine) {
	r.GET("")
}
func InitMyRouters() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	r.Run(":" + viper.GetString("port"))
}
