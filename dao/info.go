package dao

import (
	"database/sql"
	"gorm.io/gorm"
	"herbalBody/model"
)

type InfoDaoImpl struct {
	db  *sql.DB
	gdb *gorm.DB
}

func (i *InfoDaoImpl) AddInfo(info model.BasicInfo) error {
	result := i.gdb.Create(&model.BasicInfo{UserId: info.UserId})
	return result.Error
}

func (i *InfoDaoImpl) SearchInfo(userId int) (info model.BasicInfo, err error) {
	result := i.gdb.Where(&model.BasicInfo{UserId: userId}).First(&info)
	return info, result.Error
}

func (i *InfoDaoImpl) UpdateInfo(userId int, info model.BasicInfo) error {
	result := i.gdb.Take(&model.BasicInfo{}).Where(&model.BasicInfo{UserId: userId}).Updates(&info)
	return result.Error
}

func (i *InfoDaoImpl) AddBasic(basic model.BasicData) error {
	result := i.gdb.Create(&basic)
	return result.Error
}

func (i *InfoDaoImpl) SearchBasic(userId int) (basics []model.BasicData, err error) {
	result := i.gdb.Where(&model.BasicData{UserId: userId}).Find(&basics)
	return basics, result.Error
}

func (i *InfoDaoImpl) AddMotor(motor model.MotorData) error {
	result := i.gdb.Create(&motor)
	return result.Error
}

func (i *InfoDaoImpl) SearchMotor(userId int) (motors []model.MotorData, err error) {
	result := i.gdb.Where(&model.MotorData{UserId: userId}).Find(&motors)
	return motors, result.Error
}

func (i *InfoDaoImpl) AddConclusion(conclusion model.Conclusion) error {
	result := i.gdb.Create(&model.Conclusion{UserId: conclusion.UserId})
	return result.Error
}

func (i *InfoDaoImpl) SearchConclusion(userId int) (conclusion model.Conclusion, err error) {
	result := i.gdb.Where(&model.Conclusion{UserId: userId}).First(&conclusion)
	return conclusion, result.Error
}

func (i *InfoDaoImpl) UpdateConclusion(userId int, conclusion model.Conclusion) error {
	result := i.gdb.Take(&model.Conclusion{}).Where(&model.Conclusion{UserId: userId}).Updates(&conclusion)
	return result.Error
}

func NewInfoDao() *InfoDaoImpl {
	return &InfoDaoImpl{
		db:  DB,
		gdb: GDB,
	}
}
