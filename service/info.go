package service

import (
	"gorm.io/gorm"
	"herbalBody/dao"
	"herbalBody/model"
	"herbalBody/util"
)

type InfoServiceImpl struct {
	InfoDao
	*gorm.DB
}

func (i *InfoServiceImpl) SearchInfo(userId int) (code int, info model.BasicInfo) {
	info, err := i.InfoDao.SearchInfo(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.ErrRecordNotFoundCode, model.BasicInfo{}
		} else {
			return util.InternalServerErrCode, model.BasicInfo{}
		}
	}
	return util.NoErrCode, info
}

func (i *InfoServiceImpl) UpdateInfo(userId int, info model.BasicInfo) int {
	err := i.InfoDao.UpdateInfo(userId, info)
	if err != nil {
		return util.InternalServerErrCode
	}
	return util.NoErrCode
}

func (i *InfoServiceImpl) SearchBasic(userId int) (code int, basics []model.BasicData) {
	basics, err := i.InfoDao.SearchBasic(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.ErrRecordNotFoundCode, nil
		} else {
			return util.InternalServerErrCode, nil
		}
	}
	return util.NoErrCode, basics
}

func (i *InfoServiceImpl) AddBasic(basic model.BasicData) int {
	err := i.InfoDao.AddBasic(basic)
	if err != nil {
		return util.InternalServerErrCode
	}
	return util.NoErrCode
}

func (i *InfoServiceImpl) AddMotor(motor model.MotorData) int {
	err := i.InfoDao.AddMotor(motor)
	if err != nil {
		return util.InternalServerErrCode
	}
	return util.NoErrCode
}

func (i *InfoServiceImpl) SearchMotor(userId int) (code int, motors []model.MotorData) {
	motors, err := i.InfoDao.SearchMotor(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.ErrRecordNotFoundCode, nil
		} else {
			return util.InternalServerErrCode, nil
		}
	}
	return util.NoErrCode, motors
}

func (i *InfoServiceImpl) SearchConclusion(userId int) (code int, conclusion model.Conclusion) {
	conclusion, err := i.InfoDao.SearchConclusion(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.ErrRecordNotFoundCode, model.Conclusion{}
		} else {
			return util.InternalServerErrCode, model.Conclusion{}
		}
	}
	return util.NoErrCode, conclusion
}

func (i *InfoServiceImpl) UpdateConclusion(userId int, conclusion model.Conclusion) int {
	err := i.InfoDao.UpdateConclusion(userId, conclusion)
	if err != nil {
		return util.InternalServerErrCode
	}
	return util.NoErrCode
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
	SearchMotor(userId int) (motors []model.MotorData, err error)
	AddConclusion(conclusion model.Conclusion) error
	SearchConclusion(userId int) (conclusion model.Conclusion, err error)
	UpdateConclusion(userId int, conclusion model.Conclusion) error
}
