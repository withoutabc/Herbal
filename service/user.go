package service

import (
	"gorm.io/gorm"
	"herbalBody/dao"
	"herbalBody/model"
	"herbalBody/mylog"
	"herbalBody/util"
	"time"
)

type UserDaoImpl struct {
	InfoDao
	UserDao
	*gorm.DB
}

func NewUserServiceImpl() *UserDaoImpl {
	return &UserDaoImpl{
		InfoDao: dao.NewInfoDao(),
		UserDao: dao.NewUserDao(),
		DB:      dao.GetGDB(),
	}
}

type UserDao interface {
	SearchUserByUsername(username string) (user model.User, err error)
	InsertUser(u model.RegisterUser) (err error)
	InsertSignature(userId int) (err error)
	SearchUserIdByUsername(string) (int, error)
	UpdatePassword(userId int, password string) (err error)
	SearchUserById(int) (model.User, error)
}

func (u *UserDaoImpl) RegisterService(registerUser model.RegisterUser) (code int32, err error) {
	//判断两次密码是否相同
	if registerUser.Password != registerUser.ConfirmPassword {
		//101密码不相同
		return 101, nil
	}
	//校验合法性
	if len(registerUser.Username) != 11 || registerUser.Username[0] != '1' {
		//102手机号位数不符
		return 102, nil
	}
	if len(registerUser.Password) < 6 {
		//103密码长度小于6
		return 103, nil
	}
	//手机号是否重复
	mysqlUser, err := u.UserDao.SearchUserByUsername(registerUser.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		mylog.Log.Printf("search user err:%v\n", err)
		return 100, err
	}
	if mysqlUser.Password != "" {
		//104 用户名已存在
		return 104, nil
	}
	//transaction
	tx := u.DB.Begin()
	//create user
	if ResultErrorRollback(tx, tx.Create(&model.User{
		Username: registerUser.Username,
		Password: registerUser.Password,
		Role:     registerUser.Role,
	}).Error) != 0 {
		return util.TransactionErrorCode, util.TransactionError
	}
	//search user id
	var user model.User
	if ResultErrorRollback(tx, tx.Where(&model.User{Username: registerUser.Username}).First(&user).Error) != 0 {
		return util.TransactionErrorCode, util.TransactionError
	}
	//create signature
	if ResultErrorRollback(tx, tx.Create(&model.Signature{UserId: user.UserId}).Error) != 0 {
		return util.TransactionErrorCode, util.TransactionError
	}
	//create basic info
	if ResultErrorRollback(tx, tx.Create(&model.BasicInfo{UserId: user.UserId}).Error) != 0 {
		return util.TransactionErrorCode, util.TransactionError
	}
	//create conclusions
	if ResultErrorRollback(tx, tx.Create(&model.Conclusion{UserId: user.UserId}).Error) != 0 {
		return util.TransactionErrorCode, util.TransactionError
	}
	//correct
	return 0, tx.Commit().Error
}

func (u *UserDaoImpl) LoginService(user model.LoginUser) (model.Login, int32, error) {
	//查询用户名是否存在
	mysqlUser, err := u.SearchUserByUsername(user.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 101 用户名不存在
			return model.Login{}, 101, nil
		} else {
			mylog.Log.Printf("search password err:%v\n", err)
			return model.Login{}, 100, nil
		}
	}
	//验证密码是否正确
	if mysqlUser.Password != user.Password {
		// 102 密码错误
		return model.Login{}, 102, nil
	}
	//成功登录，设置token
	accessToken, refreshToken, claims, err := GenToken(mysqlUser.UserId, mysqlUser.Role)
	if err != nil {
		mylog.Log.Printf("gen token err:%v\n", err)
		return model.Login{}, 100, nil
	}
	return model.Login{
		UserId:       claims.UserId,
		Role:         claims.Role,
		LoginTime:    time.Now(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, 0, nil
}

func (u *UserDaoImpl) RefreshTokenService(token model.RefreshToken) (model.Login, int32, error) {
	accessToken, refreshToken, claims, err := RefreshToken(token.RefreshToken)
	if err != nil {
		if err.Error() == "错误的类型" {
			//101token类型错误
			return model.Login{}, 101, nil
		}
		if err.Error() == "invalid refresh token signature" {
			//102签名认证错误
			return model.Login{}, 102, nil
		}
		mylog.Log.Printf("refresh token err:%v\n", err)
		return model.Login{}, 100, nil
	}
	return model.Login{
		UserId:       claims.UserId,
		Role:         claims.Role,
		LoginTime:    time.Now(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, 0, nil
}

func (u *UserDaoImpl) ChangePassword(userId int, rcp model.ReqChangePwd) (code int) {
	//check password
	user, err := u.UserDao.SearchUserById(userId)
	if err != nil {
		return util.InternalServerErrCode
	}
	if user.Password != rcp.OldPassword {
		return util.WrongPasswordCode
	}
	//update password
	err = u.UserDao.UpdatePassword(userId, rcp.NewPassword)
	if err != nil {
		if err == util.ErrRowsAffected {
			return util.ErrRowsAffectedCode
		}
		return util.InternalServerErrCode
	}
	return util.NoErrCode
}

func SearchUsernameByUserId(userId int) (err error, username string) {
	err, username = dao.SearchUsernameByUserId(userId)
	return err, username
}
