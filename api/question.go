package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/mylog"
	"herbalBody/service"
	"herbalBody/util"
	"net/http"
)

func GetQuestionnaire(c *gin.Context) {
	questionnaires, err := service.Query()
	if err != nil {
		util.RespInternalErr(c)
		mylog.Log.Printf("query err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespQuestionnaire{
		Status: 200,
		Data:   questionnaires,
	})
}
