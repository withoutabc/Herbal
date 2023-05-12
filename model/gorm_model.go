package model

type Questions struct {
	QuestionnaireId int    `json:"questionnaire_id"` //问卷id,九种体质各对应一个id，1-9
	QuestionId      int    `json:"question_id"`      //问题id
	Question        string `json:"question" `        //问题内容
}

type Option struct {
	QuestionnaireId int    `json:"questionnaire_id"` //问卷id
	QuestionId      int    `json:"question_id"`      //问题id
	OptionId        int    `json:"option_id"`        //选项id
	Option          string `json:"option"`           //选项内容
}

type User struct {
	UserId   int    `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"role"`
	Role     string `json:"role" gorm:"not null"` // 角色共三种，详细看roleAuth.go
}

// Questionnaires 九种体质id为1-9
type Questionnaires struct {
	QuestionnaireId int    `json:"questionnaire_id"` //问题id
	QuestionId      int    `json:"question_id"`      //每种体质的问题id
	Suggestion      string `json:"suggestion"`       //该体质的建议
}

// Grade 对应每个用户九种体质的成绩
type Grade struct {
	UserId int `json:"user_id"` //用户id
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
