package model

import "time"

type Login struct {
	UserId       int       `json:"user_id"`
	LoginTime    time.Time `json:"login_time"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RespLogin struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Login  `json:"data"`
}

type RespToken struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Login  `json:"data"`
}

type RespSubmission struct {
	Status int         `json:"status" form:"status"`
	Data   WholeResult `json:"data" form:"data" `
}

type WholeResult struct {
	QuestionnaireId int      `json:"questionnaire_id" form:"questionnaire_id"`
	Questions       []Result `json:"questions" form:"questions"`
}

type Result struct {
	QuestionId int `json:"question_id" form:"question_id"`
	OptionId   int `json:"option_id" form:"option_id"`
}
