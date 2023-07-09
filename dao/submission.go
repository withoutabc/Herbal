package dao

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"herbalBody/model"
	"herbalBody/mylog"
	"sort"
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

type SubmissionDao struct {
	db  *sql.DB
	gdb *gorm.DB
}

func (q *SubmissionDao) QueryQuestionnaire(questionnaireId int) (questionnaire string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewSubmissionDao() *SubmissionDao {
	return &SubmissionDao{
		db:  DB,
		gdb: GDB,
	}
}

func QueryIfExistQuestion(s model.Submission, questionId int) (err error, b bool) {
	var answer string
	err = DB.QueryRow("select answer_id from submission where user_id=? and step=? and question_id=?", s.UserId, s.Step, questionId).Scan(&answer)
	if err != nil && err != sql.ErrNoRows {
		mylog.Log.Println("")
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
	_, err = DB.Exec("insert into grades (user_id) values (?)", userId)
	return err
}

func QueryIfGrade(userId int) (err error, b bool) {
	var a, k, c, d, e, f, g, h, i, j int
	err = DB.QueryRow("select * from grades where user_id=?", userId).Scan(&a, &c, &d, &e, &f, &g, &h, &i, &j, &k)
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
	sqL.WriteString("update grades set ")
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
	err = DB.QueryRow("select `option` from `options` where questionnaire_id=? and question_id=? and option_id=?", questionnaireId, questionId, answerId).Scan(&answer)
	if err != nil {
		return "", err
	}
	return answer, nil
}

func (q *SubmissionDao) IfExistQuestionId(questionnaireId int, questionId int) (exist bool, err error) {
	var question model.Questions
	result := q.gdb.Where(&model.Questions{
		QuestionnaireId: questionnaireId,
		QuestionId:      questionId,
	}).Find(&question)
	if question.Question == "" {
		return false, result.Error
	}
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (q *SubmissionDao) FindMaxQuestionId(questionnaireId int) (max int, err error) {
	var questions []model.Questions
	result := q.gdb.Last(&model.Questions{
		QuestionnaireId: questionnaireId,
	}).Find(&questions)
	if result.Error != nil {
		return 0, result.Error
	}
	var maxes []int
	for _, question := range questions {
		maxes = append(maxes, question.QuestionId)
	}
	sort.Ints(maxes)
	return maxes[len(maxes)-1], result.Error
}

func (q *SubmissionDao) FindAllAnswers(questionnaireId int, questionId int) (answers []int, err error) {
	var options []model.Option
	results := q.gdb.Where(&model.Option{
		QuestionnaireId: questionnaireId,
		QuestionId:      questionId,
	}).Find(&options)
	for _, option := range options {
		answers = append(answers, option.OptionId)
	}
	return answers, results.Error
}

func (q *SubmissionDao) QueryGrade(userId int, questionnaireId int) (grade int, err error) {
	result := q.gdb.Table("grades").Select(map1[questionnaireId]).Where(&model.Grade{UserId: userId}).Scan(&grade)
	return grade, result.Error
}

func (q *SubmissionDao) QuerySuggestion(questionnaireId int) (suggestion string, err error) {
	result := q.gdb.Where(&model.Questions{QuestionnaireId: questionnaireId}).First(&suggestion)
	return suggestion, result.Error
}

//func QueryGrade(userId int, questionnaireId int) (grade int, err error) {
//	var sqL strings.Builder
//	sqL.WriteString("select ")
//	sqL.WriteString(map1[questionnaireId])
//	sqL.WriteString(" from grade where user_id=?")
//	err = DB.QueryRow(sqL.String(), userId).Scan(&grade)
//	return grade, err
//}
