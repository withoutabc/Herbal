package dao

import "herbalBody/model"

func SearchUserByName(username string) (err error, u model.LoginUser) {
	row := DB.QueryRow("select * from user where username=?", username)
	if err = row.Err(); err != nil {
		return nil, model.LoginUser{}
	}
	err = row.Scan(&u.UserId, &u.Username, &u.Password)
	return err, u
}

func InsertUser(u model.RegisterUser) (err error) {
	_, err = DB.Exec("insert into user (username,password) values(?,?)", u.Username, u.Password)
	return err
}
