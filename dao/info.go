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
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) SearchInfo(userId int) (info model.BasicInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) UpdateInfo(userId int, info model.BasicInfo) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) AddBasic(basic model.BasicData) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) SearchBasic(userId int) (basics []model.BasicData, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) AddMotor(motor model.MotorData) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) SearchMotor(userId int) (motors []model.BasicData, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) AddConclusion(conclusion model.Conclusion) error {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) SearchConclusion(userId int) (conclusion model.Conclusion, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfoDaoImpl) UpdateConclusion(userId int, conclusion model.Conclusion) error {
	//TODO implement me
	panic("implement me")
}

func NewInfoDao() *InfoDaoImpl {
	return &InfoDaoImpl{
		db:  DB,
		gdb: GDB,
	}
}
