package dao

import (
	"database/sql"
	"herbalBody/model"
	"log"
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
