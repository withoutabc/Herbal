package service

import (
	"database/sql"
	"herbalBody/dao"
	"herbalBody/model"
)

// Query 仅仅不到50行，返回了这么多东西，牛！
func Query() (Results []model.Questionnaire, err error) {
	for questionnaireId := 1; ; questionnaireId++ {
		var questionnaire model.Questionnaire
		questionnaireString, err := dao.QueryQuestionnaire(questionnaireId)
		//找到了这个问卷主题，写入
		if questionnaireString != "" {
			questionnaire.Term = questionnaireString
			questionnaire.QuestionnaireId = questionnaireId
			//想办法赋值给questions
			questions, err := dao.QueryQuestion(questionnaireId)
			if err != nil {
				return nil, err
			}
			//questionnaire.Questions是一个空切片，没有分配任何的内存空间
			questionnaire.Questions = make([]model.Ques, len(questions))
			//不用判断questions是否为nil，因为问卷id（一定存在）此时必能找到问题
			for questionId := 1; questionId <= len(questions); questionId++ {
				questionnaire.Questions[questionId-1] = questions[questionId-1]
				//给每个问题id找到它的选项并赋值进去，for
				options, err := dao.QueryOption(questionnaireId, questionId)
				//?感觉没必要？
				if err != nil && err != sql.ErrNoRows {
					return nil, err
				}
				questionnaire.Questions[questionId-1].Options = make([]model.Op, len(options))
				for optionId := 1; optionId <= len(options); optionId++ {
					questionnaire.Questions[questionId-1].Options[optionId-1] = options[optionId-1]
				}
			}
			if err != nil {

				return nil, err
			}
		} else {
			//找不到主题循环就结束了
			break
		}
		if err != nil {
			return nil, err
		}
		Results = append(Results, questionnaire)
	}
	return Results, nil
}
