package models

import "database/sql"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Team      string `json:"team"`
}

type AnonymousUser struct {
	ID   int
	Name string
	Key  string
}

type Player struct {
	ID              int
	UserID          int
	User            User
	AnonymousUserID int
	AnonymousUser   AnonymousUser
}

type Section struct {
	ID   int
	Name string
}

type Quiz struct {
	ID             int            `json:"id"`
	QuizName       string         `json:"quiz_name"`
	StartSchedule  sql.NullString `json:"start_schedule"`
	StartTimestamp sql.NullString `json:"start_timestamp"`
	EndTimestamp   sql.NullString `json:"end_timestamp"`
	RoomCode       string         `json:"room_code"`
	RoomPassword   string         `json:"room_password"`
	QuizQuestions  []QuizQuestion `json:"quiz_questions"`
}

type AnswerOption struct {
	ID             int
	Value          string
	Correct        bool
	QuizQuestionID int
}

type QuizQuestion struct {
	ID             int
	QuizID         int
	Quiz           Quiz
	QuestionText   string
	AnswerType     string
	AnswerText     string
	AnswerNumber   int
	AnswerOpttions []AnswerOption
}

type PlayerAnswer struct {
	ID              int
	Answer          string
	PlayerID        Player
	Timestamp       string
	TimestampClient string
}

type Tournament struct {
	ID             int
	TournamentName string
}
