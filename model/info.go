package model

import "time"

// BasicInfo 基础信息
type BasicInfo struct {
	UserId   int    `json:"user_id" form:"user_id" gorm:"not null"`
	Nickname string `json:"nickname" form:"nickname" gorm:"default:''" binding:"required"` //昵称
	Name     string `json:"name" form:"name" gorm:"default:''" binding:"required"`         //姓名
	Gender   int    `json:"gender" form:"gender" gorm:"default:0" binding:"required"`      //性别 0未知，1男，2女
	Birthday string `json:"birthday" form:"birthday" gorm:"default:''" binding:"required"` //生日
	Address  string `json:"address" form:"address" gorm:";default:''" binding:"required"`  //住址
	Phone    string `json:"phone" form:"phone" gorm:"default:''" binding:"required"`       //手机号
	Email    string `json:"email" form:"email" gorm:"default:''" binding:"required"`       //邮箱
}

// BasicData 基础数据
type BasicData struct {
	UserId int       `json:"user_id" form:"user_id"`
	Time   time.Time `json:"time" form:"time" gorm:"not null"`
	Height int       `json:"height" form:"height" gorm:"default:0" binding:"required"` //身高
	Weight int       `json:"weight" form:"weight" gorm:"default:0" binding:"required"` //体重
	Fat    int       `json:"fat" form:"fat" gorm:"default:0" binding:"required"`       //体脂
}

// MotorData 监测数据
type MotorData struct {
	UserId                int       `json:"user_id" form:"user_id" gorm:"not null" `
	Time                  time.Time `json:"time" form:"time" gorm:"not null" `
	BloodPressure         int       `json:"blood_pressure" form:"blood_pressure" gorm:"default:0" binding:"required"`                   //血压
	HeartHealth           int       `json:"heart_health" form:"heart_health" gorm:"default:0" binding:"required"`                       //心脏健康
	Stress                int       `json:"stress"  form:"stress" gorm:"default:0" binding:"required"`                                  //压力
	BloodOxygenSaturation int       `json:"blood_oxygen_saturation" form:"blood_oxygen_saturation" gorm:"default:0" binding:"required"` //血氧饱和度
}

// Conclusion 测评结论
type Conclusion struct {
	UserId     int    `json:"user_id" form:"user_id" gorm:"not null"`
	Body       string `json:"body" form:"body" gorm:"default:''" binding:"required"`             //中医体质
	Depression string `json:"depression" form:"depression" gorm:"default:''" binding:"required"` //抑郁程度
}
