package api

import (
	"github.com/gin-gonic/gin"
	"herbalBody/model"
	"herbalBody/service"
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
	SearchInfo(userId int) (info model.BasicInfo, err error)
	UpdateInfo(userId int, info model.BasicInfo) error

	SearchBasic(userId int) (basics []model.BasicData, err error)
	AddMotor(motor model.MotorData) error
	SearchMotor(userId int) (motors []model.BasicData, err error)

	SearchConclusion(userId int) (conclusion model.Conclusion, err error)
	UpdateConclusion(userId int, conclusion model.Conclusion) error
}

func (i *InfoServiceImpl) SearchInfo(c *gin.Context) {

}

func (i *InfoServiceImpl) UpdateInfo(c *gin.Context) {

}

func (i *InfoServiceImpl) SearchBasic(c *gin.Context) {

}

func (i *InfoServiceImpl) SearchMotor(c *gin.Context) {

}

func (i *InfoServiceImpl) SearchConclusion(c *gin.Context) {

}
func (i *InfoServiceImpl) UpdateConclusion(c *gin.Context) {

}
