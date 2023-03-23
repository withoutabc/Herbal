package dao

import (
	"database/sql"
	"herbalBody/model"
	"log"
	"strings"
)

var map1 = map[int]string{
	1: "YangXu",
	2: "YinXU",
	3: "QiXu",
	4: "TanShi",
	5: "ShiRe",
	6: "XueYu",
	7: "TeBin",
	8: "QiYu",
	9: "PingHe",
}

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

func QuerySubmitAnswer(userId int, questionnaireId int, questionId int) (answer string, err error) {
	var answerId int
	err = DB.QueryRow("select answer_id from submission where user_id=? and step=? and question_id=?", userId, questionnaireId, questionId).Scan(&answerId)
	if err != nil {
		return "", err
	}
	err = DB.QueryRow("select `option` from `option` where questionnaire_id=? and question_id=? and option_id=?", questionnaireId, questionId, answerId).Scan(&answer)
	if err != nil {
		return "", err
	}
	return answer, nil
}

func QueryGrade(userId int, questionnaireId int) (grade int, err error) {
	var sqL strings.Builder
	sqL.WriteString("select ")
	sqL.WriteString(map1[questionnaireId])
	sqL.WriteString(" from grade where user_id=?")
	err = DB.QueryRow(sqL.String(), userId).Scan(&grade)
	return grade, err
}

func QuerySuggestion(questionnaireId int) (suggestion string, err error) {
	err = DB.QueryRow("select suggestion from questionnaire where questionnaire_id=?", questionnaireId).Scan(&suggestion)
	return suggestion, err
}
