package service

import (
	"gorm.io/gorm"
	"herbalBody/dao"
	"herbalBody/model"
	"log"
)

type PromiseServiceImpl struct {
	PromiseService
	SignatureService
}

func (p *PromiseServiceImpl) QueryPromiseService(userId, versionId int) (model.Promises, int, error) {
	var promises model.Promises
	//先确定version
	version, err := p.QueryVersion(versionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("record not found")
			return model.Promises{}, 101, err //version不存在
		}
		log.Printf("query version err:%v\n", err)
		return model.Promises{}, 100, err
	}
	promises.Version = version
	//确定是否签署
	b, err := p.PromiseService.QueryIfSubmit(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("record not found")
			return model.Promises{}, 102, err //用户未注册
		}
		log.Printf("query if submit err:%v\n", err)
		return model.Promises{}, 100, err
	}
	promises.IsSubmit = b
	//已经签署
	if b == true {
		promises.Promise = make([]model.PromisesPart, 1)
		promises.Promise[0].Title = "承诺书已签署"
		return promises, 0, nil
	}
	//未签署,找齐所有承诺书内容
	//先找title
	var titles []model.Title
	titles, err = p.PromiseService.QueryTitle(versionId)
	promises.Promise = make([]model.PromisesPart, len(titles))
	for i := 0; i < len(titles); i++ {
		promises.Promise[i].Title = titles[i].Title
		var lists []model.List
		lists, err = p.PromiseService.QueryList(versionId, titles[i].TitleId)
		promises.Promise[i].List = make([]string, len(lists))
		for j := 0; j < len(lists); j++ {
			promises.Promise[i].List[j] = lists[j].List
		}
	}
	//change false into true
	err = p.SignatureService.ChangeSignatureStatus(userId)
	if err != nil {
		log.Printf("change signature status err:%v\n", err)
		return model.Promises{}, 100, err
	}
	return promises, 0, nil
}

func NewPromiseServiceImpl() *PromiseServiceImpl {
	return &PromiseServiceImpl{
		PromiseService:   dao.NewPromiseDao(),
		SignatureService: dao.NewSignatureDao(),
	}
}

type PromiseService interface {
	QueryVersion(versionId int) (string, error)
	QueryTitle(versionId int) ([]model.Title, error)
	QueryList(versionId int, titleId int) ([]model.List, error)
	QueryIfSubmit(userId int) (bool, error)
}

type SignatureService interface {
	ChangeSignatureStatus(userId int) error
}
