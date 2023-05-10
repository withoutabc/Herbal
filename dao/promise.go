package dao

import (
	"gorm.io/gorm"
	"herbalBody/model"
)

type PromiseDao struct {
	db *gorm.DB
}

type SignatureDao struct {
	db *gorm.DB
}

func (s SignatureDao) ChangeSignatureStatus(userId int) error {
	var signature model.Signature
	result := s.db.Where(&model.Signature{UserId: userId}).First(&signature)
	if result.Error != nil {
		return result.Error
	}
	signature.PromiseStatus = 1
	result = s.db.Where(&model.Signature{UserId: userId}).Save(&signature)
	return result.Error
}

func (p *PromiseDao) QueryIfSubmit(userId int) (bool, error) {
	var signature model.Signature
	result := p.db.Where(&model.Signature{UserId: userId}).First(&signature)
	if result.Error != nil {
		return false, result.Error
	}
	if signature.PromiseStatus == 0 {
		return false, nil
	}
	return true, nil
}

func (p *PromiseDao) QueryVersion(versionId int) (string, error) {
	var version model.Version
	result := p.db.Where(&model.Version{VersionId: versionId}).First(&version)
	return version.Version, result.Error
}

func (p *PromiseDao) QueryTitle(versionId int) ([]model.Title, error) {
	var titles []model.Title
	result := p.db.Where(&model.Title{VersionId: versionId}).Find(&titles)
	return titles, result.Error
}

func (p *PromiseDao) QueryList(versionId int, titleId int) ([]model.List, error) {
	var lists []model.List
	result := p.db.Where(model.List{VersionId: versionId, TitleId: titleId}).Find(&lists)
	return lists, result.Error
}

func NewPromiseDao() *PromiseDao {
	return &PromiseDao{
		db: GDB,
	}
}

func NewSignatureDao() *SignatureDao {
	return &SignatureDao{
		db: GDB,
	}
}
