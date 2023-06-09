package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/mylog"
	"herbalBody/service"
	"herbalBody/util"
	"net/http"
	"strconv"
)

type UserServiceImpl struct {
	UserService
}

func NewUserApi() *UserServiceImpl {
	return &UserServiceImpl{
		UserService: service.NewUserServiceImpl(),
	}
}

type UserService interface {
	RegisterService(model.RegisterUser) (code int32, err error)
	LoginService(model.LoginUser) (model.Login, int32, error)
	RefreshTokenService(token model.RefreshToken) (model.Login, int32, error)
	ChangePassword(userId int, rcp model.ReqChangePwd) (code int)
}

func (u *UserServiceImpl) Register(c *gin.Context) {
	//接收数据
	var registerUser model.RegisterUser
	err := c.ShouldBind(&registerUser)
	if err != nil {
		mylog.Log.Printf("shouldbind err:%v\n", err)
		util.RespParamErr(c)
		return
	}
	code, err := u.UserService.RegisterService(registerUser)
	if err != nil {
		mylog.Log.Printf("register err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	switch code {
	case 101:
		util.NormErr(c, 1001, "密码不相同")
		return
	case 102:
		util.NormErr(c, 1002, "手机号不合法")
		return
	case 103:
		util.NormErr(c, 1003, "密码位数小于6")
		return
	case 104:
		util.NormErr(c, 1004, "用户名已存在")
		return
	}
	util.RespOK(c, "register success")
}

func (u *UserServiceImpl) Login(c *gin.Context) {
	//接收数据
	var LoginUser model.LoginUser
	err := c.ShouldBind(&LoginUser)
	if err != nil {
		util.RespParamErr(c)
		mylog.Log.Printf("shouldbind err:%v\n", err)
		return
	}
	loginModel, code, err := u.LoginService(LoginUser)
	if err != nil {
		util.RespInternalErr(c)
		mylog.Log.Printf("internal err:%v\n", err)
		return
	}
	switch code {
	case 101:
		util.NormErr(c, 1005, "用户名不存在")
		return
	case 102:
		util.NormErr(c, 1006, "密码错误")
		return
	}
	//
	c.JSON(http.StatusOK, model.RespLogin{
		Status: 200,
		Info:   "login success",
		Data:   loginModel,
	})

}

func (u *UserServiceImpl) Refresh(c *gin.Context) {
	//接收token
	var rt model.RefreshToken
	err := c.ShouldBind(&rt)
	if err != nil {
		util.RespParamErr(c)
		mylog.Log.Printf("should bind err:%v\n", err)
		return
	}
	loginModel, code, err := u.RefreshTokenService(rt)
	switch code {
	case 101:
		util.NormErr(c, 1007, "token类型错误")
		return
	case 102:
		util.NormErr(c, 1008, "签名认证错误")
		return
	}
	c.JSON(http.StatusOK, model.RespToken{
		Status: 200,
		Info:   "refresh token success",
		Data:   loginModel,
	})
}

func (u *UserServiceImpl) ChangePassword(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	var ReqChangePwd model.ReqChangePwd
	if err = c.ShouldBind(&ReqChangePwd); err != nil {
		util.RespParamErr(c)
		return
	}
	//service
	code := u.UserService.ChangePassword(IntUserId, ReqChangePwd)
	switch code {
	case util.ErrRowsAffectedCode:
		util.NormErr(c, 1001, "数据更新失败（可能是重复）")
		return
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
		return
	case util.WrongPasswordCode:
		util.NormErr(c, 1002, "密码错误")
		return
	}
	util.RespOK(c, "success")
}
