package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

func RespOK(c *gin.Context, info string) {
	c.JSON(http.StatusOK, respTemplate{
		Status: 200,
		Info:   info,
	})
}

var ParamError = respTemplate{
	Status: 400,
	Info:   "params error",
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

var InternalErr = respTemplate{
	Status: 500,
	Info:   "internal error",
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErr)
}

var UnauthorizedErr = respTemplate{
	Status: 401,
	Info:   "wrong role",
}

func RespUnauthorizedErr(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, UnauthorizedErr)
}

func NormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusOK, respTemplate{
		status,
		info,
	})
}

func RespErr(c *gin.Context, status int, err error) {
	c.JSON(http.StatusBadRequest, respTemplate{
		status,
		err.Error(),
	})
}

type NormSuccess struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Any    any    `json:"data"`
}

func RespNormSuccess(c *gin.Context, any2 any) {
	c.JSON(http.StatusOK, NormSuccess{Status: 200,
		Info: "success",
		Any:  any2,
	})
}
