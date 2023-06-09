package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/mylog"
	"herbalBody/service"
	"herbalBody/util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type SubmissionServiceImpl struct {
	SubmissionService
}

func NewSubmissionApi() SubmissionServiceImpl {
	return SubmissionServiceImpl{
		SubmissionService: service.NewSubmissionDaoImpl(),
	}
}

type SubmissionService interface {
	IfSubmissionValid(s model.Submission) (code int, err error)
	Comment(userId int) (comment model.Comment, code int, err error)
	Submit(s model.Submission) (code int, err error)
}

func (q *SubmissionServiceImpl) ReceiveSubmission(c *gin.Context) {
	//接收数据
	var s model.Submission
	err := c.ShouldBind(&s)
	if err != nil {
		util.RespParamErr(c)
		mylog.Log.Printf("shouldbind err:%v\n", err)
		return
	}
	//判断所有数据的合法性
	code, err := q.SubmissionService.IfSubmissionValid(s)
	if err != nil {
		mylog.Log.Printf("question service err:%v\n", err)
		return
	}
	switch code {
	case 100:
		util.RespInternalErr(c)
		return
	case 101:
		util.NormErr(c, 453, "答案不存在")
		return
	case 102:
		util.NormErr(c, 454, "没有按照顺序给出选项")
		return
	case 103:
		util.NormErr(c, 455, "有不存在的问题id(问题个数超出)")
		return
	}
	//保存提交结果
	_, err = q.Submit(s)
	if err != nil {
		if err.Error() == "没有对应的问卷页" {
			util.RespErr(c, 451, err)
			return
		}
		mylog.Log.Printf("submit err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c, "submit success")
}

func (q *SubmissionServiceImpl) GetComment(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	UserID, err := strconv.Atoi(userId)
	if err != nil {
		mylog.Log.Printf("strconv atoi err:%v\n", err)
		util.NormErr(c, 449, "invalid user id")
		return
	}
	comment, _, err := q.Comment(UserID)
	if err != nil {
		util.RespInternalErr(c)
		mylog.Log.Printf("get comment err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespComment{
		Status: 200,
		Info:   "get comment success",
		Data:   comment,
	})
}

func Upload(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		util.RespParamErr(c)
		return
	}
	UserID, err := strconv.Atoi(userId)
	if err != nil {
		mylog.Log.Printf("strconv atoi err:%v\n", err)
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
		mylog.Log.Printf("os.getwd err:%v\n", err)
		util.RespInternalErr(c)
		return
	}
	c.JSON(http.StatusOK, model.RespPicture{
		Status: 200,
		Info:   "upload success",
		Path:   wd + "/picture/" + newName,
	})
}
