package service

import (
	"gorm.io/gorm"
	"herbalBody/dao"
	"herbalBody/model"
)

type InfoServiceImpl struct {
	InfoDao
	*gorm.DB
}

func (i *InfoServiceImpl) AddInfo(info model.BasicInfo) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) SearchInfo(userId int) (info model.BasicInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) UpdateInfo(userId int, info model.BasicInfo) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) AddBasic(basic model.BasicData) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) SearchBasic(userId int) (basics []model.BasicData, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) AddMotor(motor model.MotorData) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) SearchMotor(userId int) (motors []model.BasicData, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) AddConclusion(conclusion model.Conclusion) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) SearchConclusion(userId int) (conclusion model.Conclusion, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoServiceImpl) UpdateConclusion(userId int, conclusion model.Conclusion) error {
	//TODO implement me
	panic("implement me")
}

func NewInfoServiceImpl() *InfoServiceImpl {
	return &InfoServiceImpl{
		InfoDao: dao.NewInfoDao(),
		DB:      dao.GetGDB(),
	}
}

type InfoDao interface {
	AddInfo(info model.BasicInfo) error
	SearchInfo(userId int) (info model.BasicInfo, err error)
	UpdateInfo(userId int, info model.BasicInfo) error
	AddBasic(basic model.BasicData) error
	SearchBasic(userId int) (basics []model.BasicData, err error)
	AddMotor(motor model.MotorData) error
	SearchMotor(userId int) (motors []model.BasicData, err error)
	AddConclusion(conclusion model.Conclusion) error
	SearchConclusion(userId int) (conclusion model.Conclusion, err error)
	UpdateConclusion(userId int, conclusion model.Conclusion) error
}
