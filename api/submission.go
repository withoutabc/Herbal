package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
		if err.Error() == "没有对应的问卷页" {
			util.RespErr(c, 451, err)
			return
		}
		log.Printf("submit err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "submit success")
}

func Upload(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	UserID, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("strconv atoi err:%v\n", err)
		util.NormErr(c, 449, "invalid user id")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		util.RespErr(c, 450, err)
		return
	}
	err, username := service.SearchUsernameByUserId(UserID)
	ext := filepath.Ext(file.Filename)
	newName := username + ext
	// 保存文件
	err = c.SaveUploadedFile(file, "./picture/"+newName)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("os.getwd err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespPicture{
		Status: 200,
		Info:   "upload success",
		Path:   wd + "/picture/" + newName,
	})
}

func Comment(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	UserID, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("strconv atoi err:%v\n", err)
		util.NormErr(c, 449, "invalid user id")
		return
	}
	comment, err := service.Comment(UserID)
	if err != nil {
		util.RespInternalErr(c)
		log.Printf("get comment err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespComment{
		Status: 200,
		Info:   "get comment success",
		Data:   comment,
	})
}
