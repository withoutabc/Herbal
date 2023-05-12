package dao2

import (
	"herbalBody/model"
	"herbalBody/service"
	"log"
	"time"
)

func InsertCommonUser(phoneNum string) (model.RespLogin, error) {
	//先注册,开启事务
	tx := GDB.Begin()
	var user = model.User{
		Username: phoneNum,
		Password: "0",
		Role:     "common",
	}
	result := tx.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		log.Printf("create user err : %v", result.Error)
		return model.RespLogin{}, result.Error
	}

	//查询改用户信息
	var MysqlUser model.User
	result = GDB.Model(model.User{Username: phoneNum}).First(&MysqlUser)
	if result.Error != nil {
		tx.Rollback()
		log.Printf("search user err : %v", result.Error)
		return model.RespLogin{}, result.Error
	}

	//成功登录，设置token
	accessToken, refreshToken, _, err := service.GenToken(MysqlUser.UserId, MysqlUser.Role)
	if err != nil {
		tx.Rollback()
		log.Printf("gen token err:%v\n", err)
		return model.RespLogin{}, err
	}

	//写respLogin信息
	var respLogin = model.RespLogin{
		Status: 200,
		Info:   "auto register success",
		Data: model.Login{
			UserId:       MysqlUser.UserId,
			Role:         MysqlUser.Role,
			LoginTime:    time.Now(),
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return respLogin, tx.Commit().Error
}
