package service

import (
	"errors"
	"herbalBody/dao"
	"herbalBody/model"
	"log"
)

// TransInsertSubmission 实现提交答案功能
func TransInsertSubmission(s model.Submission) (err error) {
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
	_ = tx.Commit()
	log.Println("transaction success")
	return nil
}

func CountGrade(step int) {

}
