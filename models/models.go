package models

// Survey object contains lots of questions
type Survey struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

// A Question's choice
type Choice struct {
	Id    string `json:"id"`
	Text  string `json:"text"`
	Value int    `json:"value"`
}

// Question object holds a question with several choices
type Question struct {
	Id      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Choices []Choice `json:"choices"`
}

// Answer Data
type Answer struct {
	QuestionId string `json:"questionId"`
	Answer     string `json:"answer"`
}

// Result data for a given survey item and question
type Result struct {
	SurveyID   string `bson:"surveyId" json:"surveyId"`
	QuestionID string `bson:"questionId" json:"questionId"`
	Answer     string `bson:"answer" json:"answer"`
	Count      int    `bson:"count" json:"count"`
}
