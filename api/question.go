package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"log"
	"net/http"
)

func GetQuestionnaire(c *gin.Context) {
	questionnaires, err := service.Query()
	if err != nil {
		util.RespInternalErr(c)
		log.Printf("query err:%v\n", err)
		return
	}
	c.JSON(http.StatusOK, model.RespQuestionnaire{
		Status: 200,
		Data:   questionnaires,
	})
}
