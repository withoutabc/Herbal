package api2

import (
	"github.com/gin-gonic/gin"
	"herbalBody/anotherone/util2/codes"
	"herbalBody/anotherone/util2/errutil"
	"net/http"
)

func handleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(errutil.CodeError); ok {
		codeAndError(c, e)
		return true
	}
	codeAndError(c, errutil.NewWithCode(codes.ErrUnknown))
	return true
}

func jsonError(c *gin.Context, err error) {

}

func jsonMessage(c *gin.Context, err error) {

}

func codeAndError(c *gin.Context, err error) {
	if e, ok := err.(errutil.CodeError); ok {
		c.JSON(http.StatusOK, gin.H{
			"status": e.Code,
			"msg":    e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": codes.ErrUnknown,
		"msg":    err.Error(),
	})
}

func jsonData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status": codes.OK,
		"msg":    "获取数据成功",
		"data":   data,
	})
}

func jsonSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": codes.OK,
		"msg":    "操作成功",
	})
}
