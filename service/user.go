package service

import (
	"herbalBody/dao"
	"herbalBody/model"
)

func SearchUserByName(username string) (err error, u model.LoginUser) {
	err, u = dao.SearchUserByName(username)
	return err, u
}

func InsertUser(u model.RegisterUser) (err error) {
	err = dao.InsertUser(u)
	return err
}
