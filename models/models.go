package models

type Survey struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

type Choice struct {
	Id    string `json:"id"`
	Text  string `json:"text"`
	Value int    `json:"value"`
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
	SurveyID   string `bson:"surveyId" json:"surveyId"`
	QuestionID string `bson:"questionId" json:"questionId"`
	Answer     string `bson:"answer" json:"answer"`
	Count      int    `bson:"count" json:"count"`
}
