package api2

import (
	"github.com/gin-gonic/gin"
	"herbalBody/anotherone/conf"
	"herbalBody/anotherone/mylog"
	"herbalBody/anotherone/service2"
	"herbalBody/anotherone/util2/codes"
	"herbalBody/anotherone/util2/errutil"
)

var log2 = mylog.Log

// OpenIDLogin 登录接口
func OpenIDLogin(c *gin.Context) {
	query := struct {
		Iv        string `query:"iv" binding:"required"`
		PhoneData string `query:"phoneData" binding:"required"`
		Code      string `query:"code" binding:"required"`
	}{}
	if handleError(c, errutil.ToCodeError(codes.ErrGinBindingQuery, c.ShouldBindQuery(&query))) {
		log2.Error(codes.CodeErrorMap[codes.ErrGinBindingQuery])
		return
	}
	if handleError(c, service2.Auth(query.Code, query.PhoneData, query.Iv)) {
		return
	}
	jsonSuccess(c)
}

// CheckPhoneTest 临时用来查看揭秘出来的手机号
// todo 用完待删
func CheckPhoneTest(c *gin.Context) {
	jsonData(c, gin.H{
		"phone": service2.PhoneNum,
	})
}

func getAppIdAndAppSecret(c *gin.Context) {
	jsonData(c, gin.H{
		"app_id":     conf.AppID,
		"app_secret": conf.AppSecret,
	})
}
