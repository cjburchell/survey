package models

type Survey struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type Choice struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type Question struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Choices []Choice `json:"choices"`
}

type Answer struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
}

type Result struct {
	SurveyId   string `bison:"surveyId" json:"surveyId"`
	QuestionId string `bison:"questionId" json:"questionId"`
	Answer     string `bison:"answer" json:"answer"`
	Count      int    `bison:"count" json:"count"`
}
