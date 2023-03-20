package dao

import (
	"herbalBody/model"
)

// QueryOption 根据问卷id和问题id找选项
func QueryOption(questionnaireId, questionId int) (options []model.Op, err error) {
	rows, err := DB.Query("select `option` ,  option_id from `option`where questionnaire_id=? and question_id=?", questionnaireId, questionId)
	if err != nil || rows.Err() != nil {
		return nil, err
	}
	for rows.Next() {
		var option model.Op
		if err = rows.Scan(&option.Option, &option.OptionId); err != nil {
			return nil, err
		}
		options = append(options, option)
	}
	return options, nil
}

// QueryQuestion 根据问卷id找问题
func QueryQuestion(questionnaireId int) (questions []model.Ques, err error) {
	rows, err := DB.Query("select question , question_id from question where questionnaire_id=?", questionnaireId)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		var ques model.Ques
		if err = rows.Scan(&ques.Question, &ques.QuestionId); err != nil {
			return nil, err
		}
		questions = append(questions, ques)
	}
	return questions, nil
}

func QueryQuestionnaire(questionnaireId int) (questionnaire string, err error) {
	err = DB.QueryRow("select questionnaire from questionnaire where questionnaire_id=?", questionnaireId).Scan(&questionnaire)
	if err != nil {
		return "", err
	}
	return questionnaire, err
}
