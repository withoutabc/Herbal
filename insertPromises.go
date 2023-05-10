package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
	"log"
)

var DB *gorm.DB

func ConnectGorm() {
	dsn := "root:224488@tcp(127.0.0.1:3306)/herbal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	log.Println("连接成功")

}

func InsertVersion() {
	DB.Create(&model.Version{Version: "2023-04-29"})
}

func InsertPromise() {
	DB.Create(&model.Title{VersionId: 1, Title: "1.开发者处理的信息"})
	DB.Create(&model.Title{VersionId: 1, Title: "2.你的权益"})
	DB.Create(&model.Title{VersionId: 1, Title: "3.开发者对信息的存储"})
	DB.Create(&model.Title{VersionId: 1, Title: "4.信息的使用规则"})
	DB.Create(&model.Title{VersionId: 1, Title: "5.信息对外提供"})
	DB.Create(&model.Title{VersionId: 1, Title: "6.你认为开发者未遵守上述约定，或有其他的投诉建议、或未成年人个人信息保护相关问题，可通过以下方式与开发者联系;或者向微信进行投诉。邮箱: itcmad@163.com"})
}

func InsertList() {
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "根据法律规定，开发者仅处理实现小程序功能所必要的信息。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "为了简化个人中心信息完善流程，开发者将在获取你的明示同意后，收集你的微信昵称、头像。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "为了用于小程序账号注册，开发者将在获取你的明示同意后，收集你的手机号。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "为了监测用户健康状态数据，开发者将在获取你的明示同意后，收集你的微信运动步数。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你的身份证号码，用于互联网医疗问诊实名认证。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你选中的照片或视频信息，用于对用户进行网上面诊。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你的发票信息，用于合作医院机构开具缴费记录证明。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你的订单信息，用于回溯用户问诊服务、产品消费等资金流动情况。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你的操作日志，用于优化用户体验与软件系统维护"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 读取你的剪切板，用于便于用户进行快捷输入操作。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你选中的文件，用于网上问诊时上传相关照片、病历。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 1, List: "开发者 收集你的所关注账号，用于优化用户问诊偏好推荐。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 2, List: "2.1关于你的个人信息，你可以通过以下方式与开发者联系，行使查阅、复制、更正、删除等法定权利，邮箱: itcmad@163.com。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 3, List: "3.1固定存储期限五年。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 4, List: "4.1开发者将会在本指引所明示的用途内使用收集的信息4.2如开发者使用你的信息超出本指引目的或合理范围，开发者必须在变更使用目的或范围前，再次以弹窗提醒方式告知并征得你的明示同意。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 5, List: "5.1开发者承诺，不会主动共享或转让你的信息至任何第三方，如存在确需共享或转让时，开发者应当直接征得或确认第三方征得你的单独同意。"})
	DB.Create(&model.List{VersionId: 1, TitleId: 5, List: "5.2开发者承诺，不会对外公开披露你的信息，如必须公开披露时，开发者应当向你告知公开披露的目的、披露信息的类型及可能涉及的信息，并征得你的单独同意。"})
}

func main() {
	ConnectGorm()
	InsertVersion()
	InsertPromise()
	InsertList()
}
