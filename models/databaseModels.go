package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Question struct {
	gorm.Model
	Question   string     `json:"question"`
	TestCases  []TestCase `json:"testCases"`
	Difficulty string     `json:"difficulty"`
}

type TestCase struct {
	ID             uint   `json:"id"`
	QuestionID     uint   `json:"question_id"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
}
