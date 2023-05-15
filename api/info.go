package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
	"herbalBody/util"
	"strconv"
	"time"
)

type InfoServiceImpl struct {
	InfoService
}

func NewInfoApi() *InfoServiceImpl {
	return &InfoServiceImpl{
		InfoService: service.NewInfoServiceImpl(),
	}
}

type InfoService interface {
	SearchInfo(userId int) (code int, info model.BasicInfo)
	UpdateInfo(userId int, info model.BasicInfo) int

	AddBasic(basicData model.BasicData) int
	SearchBasic(userId int) (code int, basics []model.BasicData)

	AddMotor(motorData model.MotorData) int
	SearchMotor(userId int) (code int, motors []model.MotorData)

	SearchConclusion(userId int) (code int, conclusion model.Conclusion)
	UpdateConclusion(userId int, conclusion model.Conclusion) int
}

func (i *InfoServiceImpl) SearchInfo(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	//service
	code, info := i.InfoService.SearchInfo(IntUserId)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	case util.ErrRecordNotFoundCode:
		util.NormErr(c, 1002, "no record")
	}
	util.RespNormSuccess(c, info)
}

func (i *InfoServiceImpl) UpdateInfo(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	var info model.BasicInfo
	if err = c.ShouldBind(&info); err != nil {
		util.RespParamErr(c)
	}
	//service
	code := i.InfoService.UpdateInfo(IntUserId, info)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	}
	util.RespOK(c, "success")
}

func (i *InfoServiceImpl) AddBasic(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	var basic model.BasicData
	if err = c.ShouldBind(&basic); err != nil {
		util.RespParamErr(c)
	}
	basic.UserId = IntUserId
	basic.Time = time.Now()
	//service
	code := i.InfoService.AddBasic(basic)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	}
	util.RespOK(c, "success")
}

func (i *InfoServiceImpl) SearchBasic(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	//service
	code, basics := i.InfoService.SearchBasic(IntUserId)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	case util.ErrRecordNotFoundCode:
		util.NormErr(c, 1002, "no record")
	}
	util.RespNormSuccess(c, basics)
}

func (i *InfoServiceImpl) AddMotor(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	var motor model.MotorData
	if err = c.ShouldBind(&motor); err != nil {
		util.RespParamErr(c)
	}
	motor.UserId = IntUserId
	motor.Time = time.Now()
	//service
	code := i.InfoService.AddMotor(motor)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	}
	util.RespOK(c, "success")
}

func (i *InfoServiceImpl) SearchMotor(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	//service
	code, motors := i.InfoService.SearchMotor(IntUserId)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	case util.ErrRecordNotFoundCode:
		util.NormErr(c, 1002, "no record")
	}
	util.RespNormSuccess(c, motors)
}

func (i *InfoServiceImpl) SearchConclusion(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	//service
	code, conclusion := i.InfoService.SearchConclusion(IntUserId)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	case util.ErrRecordNotFoundCode:
		util.NormErr(c, 1002, "no record")
	}
	util.RespNormSuccess(c, conclusion)
}

func (i *InfoServiceImpl) UpdateConclusion(c *gin.Context) {
	//receive
	userId := c.Param("user_id")
	IntUserId, err := strconv.Atoi(userId)
	if err != nil {
		util.NormErr(c, 1001, "user id is not a int")
		return
	}
	var conclusion model.Conclusion
	if err = c.ShouldBind(&conclusion); err != nil {
		util.RespParamErr(c)
	}
	//service
	code := i.InfoService.UpdateConclusion(IntUserId, conclusion)
	switch code {
	case util.InternalServerErrCode:
		util.RespInternalErr(c)
	}
	util.RespOK(c, "success")
}
