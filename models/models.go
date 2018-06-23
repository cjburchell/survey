package models

type Question struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Choices []string `json:"choices"`
}

type SurveyAnswer struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
}

type SurveyResult struct {
	QuestionId string `bison:"questionId" json:"questionId"`
	Answer     string `bison:"answer" json:"answer"`
	Count      int    `bison:"count" json:"count"`
}
