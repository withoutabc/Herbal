package service

import (
	"errors"
	"herbalBody/dao"
	"herbalBody/model"
	"log"
	"math"
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
