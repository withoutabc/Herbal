package model

type Submission struct {
	SubmissionId int    `json:"submission_id" form:"submission_id"`
	UserId       int    `json:"user_id" form:"user_id" binding:"required"`
	Step         int    `json:"step" form:"step" binding:"required"` //step = questionnaireId
	Data         []Test `json:"data" form:"data" binding:"required"`
}

type Test struct {
	QuestionId int `json:"question_id" form:"question_id" binding:"required"` //问题id
	AnswerId   int `json:"answer_id" form:"answer_id" binding:"required"`     //答案id==>对应选项的id
}

type Comment struct {
	Result     []string `json:"result"`     //判断是什么体质
	Suggestion []string `json:"suggestion"` //给出的相应建议
}

type RespComment struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   Comment `json:"data"`
}
