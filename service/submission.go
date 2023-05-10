package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"herbalBody/dao"
	"herbalBody/model"
	"log"
	"math"
	"os"
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

var map2 = map[int]int{
	1: 5,
	2: 4,
	3: 3,
	4: 2,
	5: 1,
}

var map3 = map[int]string{
	1: "阳虚质",
	2: "阴虚质",
	3: "气虚质",
	4: "痰湿质",
	5: "湿热质",
	6: "血瘀质",
	7: "特禀质",
	8: "气郁质",
	9: "平和质",
}

type SubmissionDaoImpl struct {
	QuestionDao
	GradeDao
	SubmissionDao
}

func NewSubmissionDaoImpl() *SubmissionDaoImpl {
	return &SubmissionDaoImpl{
		QuestionDao:   dao.NewSubmissionDao(),
		GradeDao:      dao.NewSubmissionDao(),
		SubmissionDao: dao.NewSubmissionDao(),
	}
}

type QuestionDao interface {
	IfExistQuestionId(questionnaireId int, questionId int) (exist bool, err error)
	FindMaxQuestionId(questionnaireId int) (max int, err error)
	FindAllAnswers(questionnaireId int, questionId int) (answers []int, err error)
	QuerySuggestion(questionnaireId int) (suggestion string, err error)
	QueryQuestionnaire(questionnaireId int) (questionnaire string, err error)
}

type SubmissionDao interface {
}

type GradeDao interface {
	QueryGrade(userId int, questionnaireId int) (grade int, err error)
}

func (q *SubmissionDaoImpl) IfSubmissionValid(s model.Submission) (code int, err error) {
	questionnaireId := s.Step
	//查找最大的问题id
	max, err := q.FindMaxQuestionId(questionnaireId)
	if err != nil {
		log.Printf("find max question id err:%v\n", err)
		return 100, err
	}
	//判断是不是大于的id没有
	for k, test := range s.Data {
		if k+1 > max || test.QuestionId > max {
			// 103 提交的问题id不存在
			return 103, err
		}
	}
	//判断是不是每个id都有
	for k, test := range s.Data {
		//直接判断是否按照顺序，且答案是否是题目的选项
		if test.QuestionId == k+1 {
			answers, err := q.FindAllAnswers(questionnaireId, test.QuestionId)
			if err != nil {
				log.Printf("find all answers err:%v\n", err)
				return 100, err
			}

			var count = 0
			for _, answer := range answers {
				if answer == test.AnswerId {
					count = 1
				}
			}
			if count == 0 {
				//101 答案不存在
				return 101, nil
			}
		} else {
			//102 没有按照顺序给出选项
			return 102, err
		}
	}
	return 0, nil
}

// Submit 实现提交答案功能
func (q *SubmissionDaoImpl) Submit(s model.Submission) (code int, err error) {
	//插入分数
	//判断是否存在那一行数据
	err, b := dao.QueryIfGrade(s.UserId)
	if err != nil {
		log.Printf("query grade err:%v\n", err)
		return 100, err
	}
	if b == false {
		err = dao.InsertGrade(s.UserId)
		if err != nil {
			log.Printf("insert grade err:%v", err)
			return 100, err
		}
	}
	DB := dao.GetDB()
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		log.Printf("begin trans action failed! err:%v", err)
		return 100, err
	}
	sql1 := "insert into submission (user_id, step, question_id, answer_id) values (?,?,?,?)"
	sql2 := "update submission set answer_id=? where question_id=? and user_id=? and step=?"
	//预处理
	stmt1, err := tx.Prepare(sql1)
	if err != nil {
		log.Printf("prepare err:%v\n", err)
		return 100, err
	}
	stmt2, err := tx.Prepare(sql2)
	if err != nil {
		log.Printf("prepare err:%v\n", err)
		return
	}
	//for循环添加每一条答案
	for i := 0; i < len(s.Data); i++ {
		//判断用户的这个答案是否已经填写，如果已经填写则进行修改
		err, b := dao.QueryIfExistQuestion(s, s.Data[i].QuestionId)
		if err != nil {
			log.Printf("query err:%v\n", err)
			return 100, err
		}
		if b == true {
			result, err := stmt2.Exec(s.Data[i].AnswerId, s.Data[i].QuestionId, s.UserId, s.Step)
			if err != nil {
				_ = tx.Rollback()
				log.Printf("exec failed err:%v\n", err)
				return 100, err
			}
			_, err = result.RowsAffected()
			if err != nil {
				_ = tx.Rollback()
				log.Printf("exec rowsAffected err:%v\n", err)
				return 100, err
			}
			continue //重开
		}
		//添加这条数据
		result, err := stmt1.Exec(s.UserId, s.Step, s.Data[i].QuestionId, s.Data[i].AnswerId)
		if err != nil {
			_ = tx.Rollback()
			log.Printf("exec failed err:%v\n", err)
			return 100, err
		}
		n, err := result.RowsAffected()
		if err != nil {
			_ = tx.Rollback()
			log.Printf("exec rowsAffected err:%v\n", err)
			return 100, err
		}
		if n != 1 {
			_ = tx.Rollback()
			log.Printf("transaction commit error, rollback\n")
			return 100, errors.New("事务提交错误")
		}
	}
	//计算分数
	var grade int
	switch s.Step {
	case 1, 2, 3, 4, 5, 6, 7, 8:
		for _, test := range s.Data {
			grade += test.AnswerId
		}
	case 9:
		for _, test := range s.Data {
			switch test.QuestionId {
			case 1, 5, 6:
				grade += test.AnswerId
			case 2, 3, 4, 7, 8:
				grade += map2[test.AnswerId]
			}
		}
	default:
		//101没有对应的问卷页
		return 101, errors.New("没有对应的问卷页")
	}
	//转化分数
	changeGrade := float64(grade-len(s.Data)) / float64(len(s.Data)*4) * 100
	//写入转化分数
	err = dao.UpdateChangeGrade(s.UserId, map1[s.Step], int(math.Round(changeGrade)))
	if err != nil {
		_ = tx.Rollback()
		log.Printf("update grade err:%v\n", err)
		return 100, err
	}
	_ = tx.Commit()
	log.Println("transaction success")
	return 0, nil
}

// Comment 获取评价
func (q *SubmissionDaoImpl) Comment(userId int) (comment model.Comment, code int, err error) {
	//声明数组 储存成绩
	var grades []int
	var exist = false
	//查找每个分数，确定结果
	for i := 0; i < 9; i++ {
		grade, err := q.QueryGrade(userId, i+1)
		if err != nil {
			log.Printf("query grade err:%v\n", err)
			return model.Comment{}, 100, err
		}
		grades = append(grades, grade)
	}
	//判断分析
	if grades[8] >= 60 {
		//有大于40的情况
		for i := 0; i < len(grades)-1; i++ {
			if grades[i] >= 40 {
				comment.Result = append(comment.Result, map3[i+1])
				suggestion, err := q.QuestionDao.QuerySuggestion(i + 1)
				if err != nil {
					log.Printf("query suggestion err:%v\n", err)
					return model.Comment{}, 100, err
				}
				comment.Suggestion = append(comment.Suggestion, suggestion)
				exist = true
			}
		}
		if exist {
			comment.Result = append(comment.Result, "是")
			return comment, 0, err
		}
		//没有大于40的情况，下面判断是否有30到39的
		//主的肯定是平和质
		comment.Result = append(comment.Result, map3[9])
		suggestion, err := q.QuestionDao.QuerySuggestion(9)
		if err != nil {
			log.Printf("query suggestion err:%v\n", err)
			return model.Comment{}, 0, err
		}
		comment.Suggestion = append(comment.Suggestion, suggestion)
		for i := 0; i < len(grades)-1; i++ {
			if grades[i] < 40 && grades[i] >= 30 {
				comment.Result = append(comment.Result, map3[i+1])
				suggestion, err = q.QuestionDao.QuerySuggestion(i + 1)
				if err != nil {
					log.Printf("query suggestion err:%v\n", err)
					return model.Comment{}, 100, err
				}
				comment.Suggestion = append(comment.Suggestion, suggestion)
				exist = true
			}
		}
		if exist {
			return comment, 0, err
		}
		//全部小于30
		return comment, 0, err
	} else if grades[8] < 60 {
		//有大于40的
		for i := 0; i < len(grades)-1; i++ {
			if grades[i] >= 40 {
				comment.Result = append(comment.Result, map3[i])
				suggestion, err := q.QuerySuggestion(i + 1)
				if err != nil {
					log.Printf("query suggestion err:%v\n", err)
					return model.Comment{}, 100, err
				}
				comment.Suggestion = append(comment.Suggestion, suggestion)
				log.Println(i)
				exist = true
			}
		}
		if exist {
			comment.Result = append(comment.Result, "是")
			return comment, 100, err
		}
		//都小于40，有30-39的
		for i := 0; i < len(grades)-1; i++ {
			if grades[i] < 40 && grades[i] >= 30 {
				comment.Result = append(comment.Result, map3[i+1])
				suggestion, err := q.QuerySuggestion(i + 1)
				if err != nil {
					log.Printf("query suggestion err:%v\n", err)
					return model.Comment{}, 100, err
				}
				comment.Suggestion = append(comment.Suggestion, suggestion)
				exist = true
			}
		}
		if exist {
			comment.Result = append(comment.Result, "倾向是")
			return comment, 100, err
		}
		//都小于30
		comment.Result = append(comment.Result, map3[9])
		suggestion, err := q.QuerySuggestion(9)
		if err != nil {
			log.Printf("query suggestion err:%v\n", err)
			return model.Comment{}, 100, err
		}
		comment.Suggestion = append(comment.Suggestion, suggestion)
		return comment, 0, err
	} else {
		return model.Comment{}, 100, errors.New("未知错误")
	}
}

// GenExcel 实现生成excel表功能
func (q *SubmissionDaoImpl) GenExcel(userId int) (filename string, code int, err error) {
	//101 用户答案未完全提交
	var record, record2 int
	var cell string
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("os.getwd err:%v\n", err)
		return "", 100, err
	}
	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	// 创建一个名为 Sheet1 的表格
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Printf("new sheet err:%v\n", err)
		return "", 100, err
	}
	// 设置默认工作薄
	f.SetActiveSheet(index)
	// 修改工作薄名称
	err = f.SetSheetName("Sheet1", "选项与得分")
	if err != nil {
		log.Printf("set sheet name err:%v\n", err)
		return "", 100, err
	}
	// 写入
	for questionnaireId := 1; ; questionnaireId++ {
		questionnaire, err := dao.QueryQuestionnaire(questionnaireId)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("query questionnaire err:%v\n", err)
			return "", 100, err
		}
		if questionnaire != "" {
			// 写入questionnaire
			cell = fmt.Sprintf("A%d", 2*questionnaireId-1)
			err = f.SetCellValue("选项与得分", cell, questionnaire)
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", 100, err
			}
			cell = fmt.Sprintf("A%d", 2*questionnaireId)
			err = f.SetCellValue("选项与得分", cell, "选项")
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", 100, err
			}
			questions, err := dao.QueryQuestion(questionnaireId)
			if err != nil && err != sql.ErrNoRows {
				log.Printf("query question err:%v\n", err)
				return "", 100, err
			}
			for k, question := range questions {
				// 写问题
				cell = fmt.Sprintf("%c%d", 'A'+k+1, 2*questionnaireId-1)
				err = f.SetCellValue("选项与得分", cell, question.Question)
				if err != nil {
					log.Printf("set cell value err:%v\n", err)
					return "", 100, err
				}
				// 写答案
				cell = fmt.Sprintf("%c%d", 'A'+k+1, 2*questionnaireId)
				// 先找一下答案
				answer, err := dao.QuerySubmitAnswer(userId, questionnaireId, question.QuestionId)
				if answer == "" {
					if questionnaireId != 5 || (questionnaireId == 5 && question.QuestionId != 6 && question.QuestionId != 7) {
						log.Println("用户答案未提交完全")
						return "", 101, errors.New("用户答案未提交完全")
					}
				}
				if err != nil && err != sql.ErrNoRows {
					log.Printf("query submit answer err:%v\n", err)
					return "", 100, err
				}
				// 写
				//if err == sql.ErrNoRows {
				//	log.Println(questionnaireId, question.QuestionId, answer)
				//}
				if answer != "" {
					err = f.SetCellValue("选项与得分", cell, answer)
					if err != nil {
						log.Printf("set cell value err:%v\n", err)
						return "", 100, err
					}
				}
				//记录一下，出循环后用
				record = 'A' + k + 1
				record2 = 2*questionnaireId - 1
			}
			cell = fmt.Sprintf("%c%d", record+1, record2)
			err = f.SetCellValue("选项与得分", cell, "得分")
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", 100, err
			}
			cell = fmt.Sprintf("%c%d", record+1, record2+1)
			//找一下得分
			grade, err := q.GradeDao.QueryGrade(userId, questionnaireId)
			if err != nil {
				log.Printf("query err:%v\n", err)
				return "", 100, err
			}
			err = f.SetCellValue("选项与得分", cell, grade)
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", 100, err
			}
		} else {
			break
		}
	}
	//设置样式
	centerStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center"}})
	if err != nil {
		log.Printf("center style err:%v\n", err)
		return "", 100, err
	}
	err = f.SetCellStyle("选项与得分", "A1", cell, centerStyle)
	if err != nil {
		log.Printf("set cell style err:%v\n", err)
		return "", 100, err
	}
	//row := []interface{}{"1", nil, 2}
	//err = f.SetSheetRow("选项与得分", "A1", &row)
	// 设置行高列宽（固定值）
	endCol := fmt.Sprintf("%c", record+1)
	err = f.SetColWidth("选项与得分", "A", "A", 10)
	if err != nil {
		log.Printf("set col1 err:%v\n", err)
		return "", 100, err
	}
	err = f.SetColWidth("选项与得分", "B", endCol, 45)
	if err != nil {
		log.Printf("set col width err:%v\n", err)
		return "", 100, err
	}
	for i := 1; i <= record2+1; i++ {
		err := f.SetRowHeight("选项与得分", i, 20)
		if err != nil {
			log.Printf("set row height err:%v\n", err)
			return "", 100, err
		}
	}
	// 保存 Excel 文件
	// 拼一下名字
	err, username := dao.SearchUsernameByUserId(userId)
	if err != nil {
		log.Printf("search username err:%v\n", err)
		return "", 100, err
	}
	filename = fmt.Sprintf("/%s-中医智慧诊疗诊断表.xlsx", username)
	if err := f.SaveAs(wd + "/excel" + filename); err != nil {
		log.Printf("save file err:%v\n", err)
		return "", 100, err
	}
	return filename, 0, nil
}
