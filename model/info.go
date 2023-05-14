package model

import "time"

// BasicInfo 基础信息
type BasicInfo struct {
	UserId   int    `json:"user_id" form:"user_id" gorm:"not null"`
	Nickname string `json:"nickname" form:"nickname" gorm:"default:''"` //昵称
	Name     string `json:"name" form:"name" gorm:"default:''"`         //姓名
	Gender   int    `json:"gender" form:"gender" gorm:"default:0"`      //性别 0未知，1男，2女
	Birthday string `json:"birthday" form:"birthday" gorm:"default:''"` //生日
	Address  string `json:"address" form:"address" gorm:";default:''"`  //住址
	Phone    string `json:"phone" form:"phone" gorm:"default:''"`       //手机号
	Email    string `json:"email" form:"email" gorm:"default:''"`       //邮箱
}

// BasicData 基础数据
type BasicData struct {
	UserId int       `json:"user_id" form:"user_id"`
	Time   time.Time `json:"time" form:"time" gorm:"not null"`
	Height int       `json:"height" form:"height" gorm:"default:0"` //身高
	Weight int       `json:"weight" form:"weight" gorm:"default:0"` //体重
	Fat    int       `json:"fat" form:"fat" gorm:"default:0"`       //体脂
}

// MotorData 监测数据
type MotorData struct {
	UserId                int       `json:"user_id" form:"user_id" gorm:"not null"`
	Time                  time.Time `json:"time" form:"time" gorm:"not null"`
	BloodPressure         int       `json:"blood_pressure" form:"blood_pressure" gorm:"default:0"`                   //血压
	HeartHealth           int       `json:"heart_health" form:"heart_health" gorm:"default:0"`                       //心脏健康
	Stress                int       `json:"stress"  form:"stress" gorm:"default:0"`                                  //压力
	BloodOxygenSaturation int       `json:"blood_oxygen_saturation" form:"blood_oxygen_saturation" gorm:"default:0"` //血氧饱和度
}

// Conclusion 测评结论
type Conclusion struct {
	UserId     int    `json:"user_id" form:"user_id" gorm:"not null"`
	Body       string `json:"body" form:"body" gorm:"default:''"`             //中医体质
	Depression string `json:"depression" form:"depression" gorm:"default:''"` //抑郁程度
}
