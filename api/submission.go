package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"log"
)

func ReceiveSubmission(c *gin.Context) {
	//接收数据
	var s model.Submission
	err := c.ShouldBind(&s)
	if err != nil {
		util.RespParamErr(c)
		log.Printf("shouldbind err:%v\n", err)
		return
	}
	//保存提交结果
	err = service.TransInsertSubmission(s)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	//返回用户填写的答案

}
