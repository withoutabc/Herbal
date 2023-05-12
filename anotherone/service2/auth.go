package service2

import (
	"encoding/json"
	"fmt"
	"herbalBody/anotherone/mylog"
	"herbalBody/anotherone/util2"
	"herbalBody/anotherone/util2/codes"
	"herbalBody/anotherone/util2/errutil"
	"herbalBody/anotherone/util2/requester"
)

var log2 = mylog.Log

//todo 下面这行用来检查是否获取成功（待删除）

var PhoneNum string

func init() {
	PhoneNum = ""
}

const (
	appID      = "wxd96aed10a5cbf5d8"
	appSecret  = "ac0ee65a4f3c8ed8a25fb0f8a0a4278c"
	requestUrl = "https://api.weixin.qq.com/sns/jscode2session"
)

type WxSessionKeyDto struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type WxPhoneDto struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

// Auth 根据前端传来的数据获取用户的手机号并且注册本地用户
func Auth(code, phoneData, iv string) (err error) {
	//发送请求code->session
	getter := requester.Getter{
		Url: requestUrl,
		Query: map[string]string{
			"appid":      appID,
			"secret":     appSecret,
			"js_code":    code,
			"grant_type": "authorization_code",
		},
		Header: nil,
	}
	bytes, err := getter.Get()
	if err != nil {
		log2.Error("GET发送出错 ", err)
		return errutil.NewWithCode(codes.ErrServerGetFail)
	}
	session := WxSessionKeyDto{}
	err = json.Unmarshal(bytes, &session)
	if err != nil {
		log2.Error("json解析出错 ", err)
		return err
	}
	fmt.Println(session)

	//AES解密 获取手机号
	decrypt, err := util2.AesDecrypt(phoneData, session.SessionKey, iv)
	if err != nil {
		log2.Error("解密数据失败", err)
		return err
	}
	var phoneDto = WxPhoneDto{}
	err = json.Unmarshal(decrypt, &phoneDto)
	if err != nil {
		log2.Error("解析手机号信息失败", err)
		return err
	}
	var phone = phoneDto.PurePhoneNumber

	//todo 下面这行用来检查是否获取成功（待删除）
	PhoneNum = phone

	//todo 用手机号注册用户获取用户信息等

	return
}
