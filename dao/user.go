package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"herbalBody/model"
	"herbalBody/util"
)

type UserDaoImpl struct {
	db  *sql.DB
	gdb *gorm.DB
}

func (u *UserDaoImpl) UpdatePassword(userId int, password string) (err error) {
	result := u.gdb.Where(&model.User{UserId: userId}).Update("password", password)
	if result.RowsAffected != 1 {
		return util.ErrRowsAffected
	}
	return result.Error
}

func (u *UserDaoImpl) SearchUserByUsername(username string) (user model.User, err error) {
	var tempUser model.User
	result := u.gdb.Where(&model.User{
		Username: username,
	}).First(&tempUser)
	return tempUser, result.Error
}

func (u *UserDaoImpl) InsertUser(registerUser model.RegisterUser) (err error) {
	result := u.gdb.Create(&model.User{
		UserId:   registerUser.UserId,
		Username: registerUser.Username,
		Password: registerUser.Password,
		Role:     registerUser.Role,
	})
	return result.Error
}

func (u *UserDaoImpl) SearchUserIdByUsername(username string) (int, error) {
	var user model.User
	result := u.gdb.Where(&model.User{Username: username}).First(&user)
	return user.UserId, result.Error
}

func (u *UserDaoImpl) InsertSignature(userId int) (err error) {
	result := u.gdb.Create(&model.Signature{UserId: userId})
	return result.Error
}

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{
		db:  DB,
		gdb: GDB,
	}
}

func SearchUsernameByUserId(userId int) (err error, username string) {
	err = DB.QueryRow("select username from users where user_id=?", userId).Scan(&username)
	return err, username
}
