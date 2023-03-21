package dao

import (
	"database/sql"
	"herbalBody/model"
	"log"
	"strings"
)

func QueryIfExistQuestion(s model.Submission, questionId int) (err error, b bool) {
	var answer string
	err = DB.QueryRow("select answer_id from submission where user_id=? and step=? and question_id=?", s.UserId, s.Step, questionId).Scan(&answer)
	if err != nil && err != sql.ErrNoRows {
		log.Println("")
		return err, false
	}
	//没找到这行==>insert处理
	if err == sql.ErrNoRows || answer == "" {
		return nil, false
	}
	//找到这行==>update处理
	return nil, true
}

func InsertGrade(userId int) (err error) {
	_, err = DB.Exec("insert into grade (user_id) values (?)", userId)
	return err
}

func QueryIfGrade(userId int) (err error, b bool) {
	var a, k, c, d, e, f, g, h, i, j int
	err = DB.QueryRow("select * from grade where user_id=?", userId).Scan(&a, &c, &d, &e, &f, &g, &h, &i, &j, &k)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		return err, true
	}
	return nil, true
}

func UpdateChangeGrade(userId int, term string, changeGrade int) (err error) {
	var sqL strings.Builder
	sqL.WriteString("update grade set ")
	sqL.WriteString(term)
	sqL.WriteString("=? where user_id=?")
	_, err = DB.Exec(sqL.String(), changeGrade, userId)
	return err
}
