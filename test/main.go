package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"time"

	"strconv"
)

func main() {
	r := gin.New()
	r.Use(LogMiddleware())
	r.GET("/GET", func(c *gin.Context) {
		c.String(200, "GET")
	})
	r.PUT("/PUT", func(c *gin.Context) {
		c.String(200, "PUT")
	})
	r.DELETE("/DELETE", func(c *gin.Context) {
		c.String(200, "DELETE")
	})
	r.POST("/POST", func(c *gin.Context) {
		c.String(200, "POST")
	})
	r.Run(":3030")
}
func main02() {
	r := gin.Default()
	r.Use(TlsHandler(8081))
	r.GET("/testtls", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"1": 1,
		})
	})
	r.RunTLS(":8081", "crt.pem", "key.pem")
}
func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			fmt.Println(err)
			return
		}

		c.Next()
	}
}

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

		fmt.Printf("[GIN] %s | %s | %d | %s | %s | %s",
			start.Format("2006-01-02 15:04:06"),
			statusColor,
			timeSub,
			clientIP,
			methodColor,
			path,
		)

	}

}

func main01() {

	fmt.Printf("\033[30m 30\033[0m1111\n")
	fmt.Printf("\033[31m 31\033[0m1111\n")
	fmt.Printf("\033[32m 32\033[0m1111\n")
	fmt.Printf("\033[33m 33\033[0m1111\n")
	fmt.Printf("\033[34m 34\033[0m1111\n")
	fmt.Printf("\033[35m 35\033[0m1111\n")
	fmt.Printf("\033[36m 36\033[0m1111\n")
	fmt.Printf("\033[37m 37\033[0m1111\n")

	fmt.Printf("\033[40m40\033[0m1111\n")
	fmt.Printf("\033[41m41\033[0m1111\n")
	fmt.Printf("\033[42m42\033[0m1111\n")
	fmt.Printf("\033[43m43\033[0m1111\n")
	fmt.Printf("\033[44m44\033[0m1111\n")
	fmt.Printf("\033[45m45\033[0m1111\n")
	fmt.Printf("\033[46m46\033[0m1111\n")
	fmt.Printf("\033[47m47\033[0m1111\n")
}
