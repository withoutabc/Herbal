package model

type Submission struct {
	SubmissionId int    `json:"submission_id" form:"submission_id"`
	UserId       int    `json:"user_id" form:"user_id" binding:"required"`
	Step         int    `json:"step" form:"step" binding:"required"`
	Data         []Test `json:"data" form:"data" binding:"required"`
}

type Test struct {
	QuestionId int `json:"question_id"`
	AnswerId   int `json:"answer_id"`
}

type Comment struct {
	Result     []string `json:"result"`
	Suggestion []string `json:"suggestion"`
}

type RespComment struct {
	Status int     `json:"status"`
	Info   string  `json:"info"`
	Data   Comment `json:"data"`
}
