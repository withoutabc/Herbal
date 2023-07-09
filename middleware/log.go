package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"herbalBody/mylog"
	"time"
)

const (
	status200 = 42
	status404 = 43
	status500 = 41

	methodPOST   = 43
	methodDELETE = 41
	methodGET    = 44
	methodPUT    = 46
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped

		// Stop timer
		end := time.Now()
		timeSub := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		//bodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		var statusColor string
		switch statusCode {
		case 200:
			statusColor = fmt.Sprintf("\033[%dm%d\033[0m", status200, statusCode)
		case 404:
			statusColor = fmt.Sprintf("\033[%dm%d\033[0m", status404, statusCode)
		case 500:
			statusColor = fmt.Sprintf("\033[%dm%d\033[0m", status500, statusCode)
		}

		var methodColor string
		switch method {
		case "GET":
			methodColor = fmt.Sprintf("\033[%dm%s\033[0m", methodGET, method)
		case "PUT":
			methodColor = fmt.Sprintf("\033[%dm%s\033[0m", methodPUT, method)
		case "DELETE":
			methodColor = fmt.Sprintf("\033[%dm%s\033[0m", methodDELETE, method)
		case "POST":
			methodColor = fmt.Sprintf("\033[%dm%s\033[0m", methodPOST, method)
		}

		_, err := fmt.Fprintf(mylog.GetOutputWriter(), "[GIN] %s | %s | %d | %s | %s | %s",
			start.Format("2006-01-02 15:04:06"),
			statusColor,
			timeSub,
			clientIP,
			methodColor,
			path,
		)
		if err != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"msg":    "日志载入失败",
				"data":   nil,
			})
			c.Abort()
			return
		}

	}

}
