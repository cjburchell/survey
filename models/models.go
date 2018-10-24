package models

// Survey object contains lots of questions
type Survey struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Questions []Question `json:"questions"`
}

// Choice object
type Choice struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Value int    `json:"value"`
}

// Question object holds a question with several choices
type Question struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`
	Choices []Choice `json:"choices"`
}

// Answer data
type Answer struct {
	QuestionID string `json:"questionId"`
	Answer     string `json:"answer"`
}

// Result data for a given survey item and question
type Result struct {
	SurveyID   string `bson:"surveyId" json:"surveyId"`
	QuestionID string `bson:"questionId" json:"questionId"`
	Answer     string `bson:"answer" json:"answer"`
	Count      int    `bson:"count" json:"count"`
}
