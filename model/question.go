package model

type RespQuestionnaire struct {
	Status int             `json:"status" form:"status"`
	Data   []Questionnaire `json:"data" form:"data "`
}

type Questionnaire struct {
	Term            string `json:"questionnaire" form:"questionnaire"`
	QuestionnaireId int    `json:"questionnaire_id" form:"questionnaire_id"`
	Questions       []Ques `json:"questions" form:"questions"`
}

type Ques struct {
	Question   string `json:"question" form:"question"`
	QuestionId int    `json:"question_id" form:"question_id"`
	Options    []Op   `json:"options" form:"options"`
}

type Op struct {
	Option   string `json:"option" form:"option"`
	OptionId int    `json:"option_id" form:"option_id"`
}
