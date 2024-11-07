package models

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RunRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Input    string `json:"input"`
}

type CreateQuestionRequest struct {
	Question  string     `json:"question"`
	TestCases []TestCase `json:"testCases"`
	Score     int        `json:"score"`
	Difficulty string    `json:"difficulty"`
}

type FetchQuestionRequest struct {
	QuestionID uint `json:"questionID"`
}
