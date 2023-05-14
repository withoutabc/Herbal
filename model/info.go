package model

import "time"

// BasicInfo 基础信息
type BasicInfo struct {
	UserId   int    `json:"user_id" form:"user_id"`
	Nickname string `json:"nickname" form:"nickname"`              //昵称
	Name     string `json:"name" form:"name"`                      //姓名
	Gender   int    `json:"gender" form:"gender" gorm:"default:0"` //性别 0未知，1男，2女
	Birthday string `json:"birthday" form:"birthday"`              //生日
	Address  string `json:"address" form:"address"`                //住址
	Phone    string `json:"phone" form:"phone"`                    //手机号
	Email    string `json:"email" form:"email"`                    //邮箱
}

// BasicData 基础数据
type BasicData struct {
	UserId int       `json:"user_id" form:"user_id"`
	Time   time.Time `json:"time" form:"time"`
	Height int       `json:"height" form:"height"` //身高
	Weight int       `json:"weight" form:"weight"` //体重
	Fat    int       `json:"fat" form:"fat"`       //体脂
}

// MotorData 监测数据
type MotorData struct {
	UserId                int       `json:"user_id" form:"user_id"`
	Time                  time.Time `json:"time" form:"time"`
	BloodPressure         int       `json:"blood_pressure" form:"blood_pressure"`                   //血压
	HeartHealth           int       `json:"heart_health" form:"heart_health"`                       //心脏健康
	Stress                int       `json:"stress"  form:"stress"`                                  //压力
	BloodOxygenSaturation int       `json:"blood_oxygen_saturation" form:"blood_oxygen_saturation"` //血氧饱和度
}

// Conclusion 测评结论
type Conclusion struct {
	UserId     int    `json:"user_id" form:"user_id"`
	Body       string `json:"body" form:"body"`             //中医体质
	Depression string `json:"depression" form:"depression"` //抑郁程度
}
