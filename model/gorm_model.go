package model

type Questions struct {
	QuestionnaireId int    `json:"questionnaire_id"`
	QuestionId      int    `json:"question_id"`
	Question        string `json:"question" `
}

type Option struct {
	QuestionnaireId int    `json:"questionnaire_id"`
	QuestionId      int    `json:"question_id"`
	OptionId        int    `json:"option_id"`
	Option          string `json:"option"`
}

type User struct {
	UserId   int    `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"role"`
	Role     string `json:"role" gorm:"not null"`
}

type Questionnaires struct {
	QuestionnaireId int    `json:"questionnaire_id"`
	QuestionId      int    `json:"question_id"`
	Suggestion      string `json:"suggestion"`
}

type Grade struct {
	UserId int `json:"user_id"`
	YangXu int `json:"yang_xu"`
	YinXu  int `json:"yin_xu"`
	QiXu   int `json:"qi_xu"`
	TanShi int `json:"tan_shi"`
	ShiRe  int `json:"shi_re"`
	XueYu  int `json:"xue_yu"`
	TeBin  int `json:"te_bin"`
	QiYu   int `json:"qi_yu"`
	PingHe int `json:"ping_he"`
}
