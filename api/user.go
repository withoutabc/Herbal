package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	//接收数据
	var u model.RegisterUser
	err := c.ShouldBind(&u)
	if err != nil {
		log.Printf("shouldbind err:%v\n", err)
		util.RespParamErr(c)
		return
	}
	//判断两次密码是否相同
	if u.Password != u.ConfirmPassword {
		util.NormErr(c, 440, "两次密码不相同")
		return
	}
	//校验合法性
	if len(u.Username) != 11 || u.Username[0] != '1' {
		//fmt.Printf("length:%d\n", len(u.Username))
		//fmt.Printf("[1]:%d\n", u.Username[0])
		util.NormErr(c, 441, "账号格式不正确")
		return
	}
	if len(u.Password) < 6 {
		util.NormErr(c, 442, "密码长度小于6")
		return
	}
	//手机号是否重复
	err, user := service.SearchUserByName(u.Username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search password err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	if user.Password != "" {
		util.NormErr(c, 443, "账户已注册")
		return
	}
	//写入数据库
	err = service.InsertUser(u)
	if err != nil {
		log.Printf("insert user err:%v\n", err)
		return
	}
	//正确返回
	util.RespOK(c, "register success")
}

func Login(c *gin.Context) {
	//接收数据
	var u model.LoginUser
	err := c.ShouldBind(&u)
	if err != nil {
		util.RespParamErr(c)
		log.Printf("shouldbind err:%v\n", err)
		return
	}
	//查询用户名是否存在
	err, user := service.SearchUserByName(u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 444, "登录：用户名不存在")
			return
		} else {
			log.Printf("search password err:%v\n", err)
			util.RespInternalErr(c)
			return
		}
	}
	//验证密码是否正确
	if user.Password != u.Password {
		util.NormErr(c, 445, "登录：密码错误")
		return
	}
	//成功登录，设置token
	accessToken, refreshToken, claims, err := service.GenToken(user.UserId)
	if err != nil {
		util.RespInternalErr(c)
		log.Printf("gen token err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespLogin{
		Status: 200,
		Info:   "login success",
		Data: model.Login{
			UserId:       user.UserId,
			LoginTime:    claims.LoginTime,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

func Refresh(c *gin.Context) {
	var rt model.RefreshToken
	err := c.ShouldBind(&rt)
	if err != nil {
		util.RespParamErr(c)
		log.Printf("should bind err:%v\n", err)
		return
	}
	accessToken, refreshToken, claims, err := service.RefreshToken(rt.RefreshToken)
	if err != nil {
		if err.Error() == "错误的类型" {
			util.RespErr(c, 451, err)
			log.Println(err)
			return
		}
		if err.Error() == "invalid refresh token signature" {
			util.RespErr(c, 452, err)
			log.Println(err)
			return
		}
		util.RespErr(c, 453, errors.New("invalid token"))
		log.Printf("refresh token err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespToken{
		Status: 200,
		Info:   "refresh token success",
		Data: model.Login{
			UserId:       claims.UserId,
			LoginTime:    claims.LoginTime,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
