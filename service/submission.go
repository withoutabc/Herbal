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

// Submit 实现提交答案和算分功能
func Submit(s model.Submission) (err error) {
	//插入分数
	//判断是否存在那一行数据
	err, b := dao.QueryIfGrade(s.UserId)
	if err != nil {
		log.Printf("query grade err:%v\n", err)
		return err
	}

	if b == false {
		err = dao.InsertGrade(s.UserId)
		if err != nil {
			log.Printf("insert grade err:%v", err)
			return err
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
		return err
	}
	sql1 := "insert into submission (user_id, step, question_id, answer_id) values (?,?,?,?)"
	sql2 := "update submission set answer_id=? where question_id=? and user_id=? and step=?"
	//预处理
	stmt1, err := tx.Prepare(sql1)
	if err != nil {
		log.Printf("prepare err:%v\n", err)
		return err
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
			return err
		}
		if b == true {
			result, err := stmt2.Exec(s.Data[i].AnswerId, s.Data[i].QuestionId, s.UserId, s.Step)
			if err != nil {
				_ = tx.Rollback()
				log.Printf("exec failed err:%v\n", err)
				return err
			}
			_, err = result.RowsAffected()
			if err != nil {
				_ = tx.Rollback()
				log.Printf("exec rowsAffected err:%v\n", err)
				return err
			}
			continue //重开
		}
		//添加这条数据
		result, err := stmt1.Exec(s.UserId, s.Step, s.Data[i].QuestionId, s.Data[i].AnswerId)
		if err != nil {
			_ = tx.Rollback()
			log.Printf("exec failed err:%v\n", err)
			return err
		}
		n, err := result.RowsAffected()
		if err != nil {
			_ = tx.Rollback()
			log.Printf("exec rowsAffected err:%v\n", err)
			return err
		}
		if n != 1 {
			_ = tx.Rollback()
			log.Printf("transaction commit error, rollback\n")
			return errors.New("事务提交错误")
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
		return errors.New("没有对应的问卷页")
	}
	//转化分数
	changeGrade := float64(grade-len(s.Data)) / float64(len(s.Data)*4) * 100
	//写入转化分数
	err = dao.UpdateChangeGrade(s.UserId, map1[s.Step], int(math.Round(changeGrade)))
	if err != nil {
		_ = tx.Rollback()
		log.Printf("update grade err:%v\n", err)
		return err
	}
	_ = tx.Commit()
	log.Println("transaction success")
	return nil
}

func GenExcel(userId int) (filename string, f *excelize.File, err error) {
	var record, record2 int
	var cell string
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("os.getwd err:%v\n", err)
		return "", nil, err
	}
	// 创建一个新的 Excel 文件
	f = excelize.NewFile()
	// 创建一个名为 Sheet1 的表格
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Printf("new sheet err:%v\n", err)
		return "", nil, err
	}
	// 设置默认工作薄
	f.SetActiveSheet(index)
	// 修改工作薄名称
	err = f.SetSheetName("Sheet1", "选项与得分")
	if err != nil {
		log.Printf("set sheet name err:%v\n", err)
		return "", nil, err
	}
	// 写入
	for questionnaireId := 1; ; questionnaireId++ {
		questionnaire, err := dao.QueryQuestionnaire(questionnaireId)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("query questionnaire err:%v\n", err)
			return "", nil, err
		}
		if questionnaire != "" {
			// 写入questionnaire
			cell = fmt.Sprintf("A%d", 2*questionnaireId-1)
			err = f.SetCellValue("选项与得分", cell, questionnaire)
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", nil, err
			}
			cell = fmt.Sprintf("A%d", 2*questionnaireId)
			err = f.SetCellValue("选项与得分", cell, "选项")
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", nil, err
			}
			questions, err := dao.QueryQuestion(questionnaireId)
			if err != nil && err != sql.ErrNoRows {
				log.Printf("query question err:%v\n", err)
				return "", nil, err
			}
			for k, question := range questions {
				// 写问题
				cell = fmt.Sprintf("%c%d", 'A'+k+1, 2*questionnaireId-1)
				err = f.SetCellValue("选项与得分", cell, question.Question)
				if err != nil {
					log.Printf("set cell value err:%v\n", err)
					return "", nil, err
				}
				// 写答案
				cell = fmt.Sprintf("%c%d", 'A'+k+1, 2*questionnaireId)
				// 先找一下答案
				answer, err := dao.QuerySubmitAnswer(userId, questionnaireId, question.QuestionId)
				if answer == "" {
					log.Println("用户答案未提交完全")
					return "", nil, errors.New("用户答案未提交完全")
				}
				if err != nil {
					log.Printf("query submit answer err:%v\n", err)
					return "", nil, err
				}
				// 写
				err = f.SetCellValue("选项与得分", cell, answer)
				if err != nil {
					log.Printf("set cell value err:%v\n", err)
					return "", nil, err
				}
				//记录一下，出循环后用
				record = 'A' + k + 1
				record2 = 2*questionnaireId - 1
			}
			cell = fmt.Sprintf("%c%d", record+1, record2)
			err = f.SetCellValue("选项与得分", cell, "得分")
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", nil, err
			}
			cell = fmt.Sprintf("%c%d", record+1, record2+1)
			//找一下得分
			grade, err := dao.QueryGrade(userId, questionnaireId)
			if err != nil {
				log.Printf("query err:%v\n", err)
				return "", nil, err
			}
			err = f.SetCellValue("选项与得分", cell, grade)
			if err != nil {
				log.Printf("set cell value err:%v\n", err)
				return "", nil, err
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
		return "", nil, err
	}
	err = f.SetCellStyle("选项与得分", "A1", cell, centerStyle)
	if err != nil {
		log.Printf("set cell style err:%v\n", err)
		return "", nil, err
	}
	//row := []interface{}{"1", nil, 2}
	//err = f.SetSheetRow("选项与得分", "A1", &row)
	// 设置行高列宽（固定值）
	endCol := fmt.Sprintf("%c", record+1)
	err = f.SetColWidth("选项与得分", "A", "A", 10)
	if err != nil {
		log.Printf("set col1 err:%v\n", err)
		return "", nil, err
	}
	err = f.SetColWidth("选项与得分", "B", endCol, 45)
	if err != nil {
		log.Printf("set col width err:%v\n", err)
		return "", nil, err
	}
	for i := 1; i <= record2+1; i++ {
		err := f.SetRowHeight("选项与得分", i, 20)
		if err != nil {
			log.Printf("set row height err:%v\n", err)
			return "", nil, err
		}
	}
	// 保存 Excel 文件
	// 拼一下名字
	err, username := dao.SearchUsernameByUserId(userId)
	if err != nil {
		log.Printf("search username err:%v\n", err)
		return "", nil, err
	}
	filename = fmt.Sprintf("/%s-中医智慧诊疗诊断表.xlsx", username)
	if err := f.SaveAs(wd + filename); err != nil {
		log.Printf("save file err:%v\n", err)
		return "", nil, err
	}
	return filename, f, nil
}
