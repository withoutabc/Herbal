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
	err = service.Submit(s)
	if err != nil {
		log.Printf("submit err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "submit success")
}
