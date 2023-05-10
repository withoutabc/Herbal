package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"net/http"
)

type PromiseImpl struct {
	PromiseApi
}

func NewPromiseApi() *PromiseImpl {
	return &PromiseImpl{
		PromiseApi: service.NewPromiseServiceImpl(),
	}
}

type PromiseApi interface {
	QueryPromiseService(userId, versionId int) (model.Promises, int, error)
}

func (p *PromiseImpl) QueryPromises(c *gin.Context) {
	var reqPromises model.ReqPromises
	err := c.ShouldBind(&reqPromises)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	promises, code, err := p.PromiseApi.QueryPromiseService(reqPromises.UserId, reqPromises.VersionId)
	switch code {
	case 100:
		util.RespInternalErr(c)
		return
	case 101:
		util.NormErr(c, 1001, "版本号不存在")
		return
	case 102:
		util.NormErr(c, 1002, "用户未注册")
		return
	}
	c.JSON(http.StatusOK, model.RespPromises{
		Status: 200,
		Info:   "query promise success",
		Data:   promises,
	})
}
