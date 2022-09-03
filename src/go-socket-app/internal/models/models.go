package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

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
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Player struct {
	ID              int           `json:"id"`
	UserID          int           `json:"user_id"`
	User            User          `json:"user"`
	AnonymousUserID int           `json:"anonymous_user_id"`
	AnonymousUser   AnonymousUser `json:"anonymous_user"`
}

type Section struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Quiz struct {
	ID             int            `json:"id"`
	QuizName       string         `json:"quiz_name"`
	StartSchedule  NullTime       `json:"start_schedule"`
	StartTimestamp NullTime       `json:"start_timestamp"`
	EndTimestamp   NullTime       `json:"end_timestamp"`
	RoomCode       string         `json:"room_code"`
	RoomPassword   NullString     `json:"room_password"`
	OrganizerId    int            `json:"organizer_id"`
	QuizState      string         `json:"quiz_state"`
	QuizQuestion   int            `json:"quiz_question"`
	QuestionState  int            `json:"question_state"`
	QuizQuestions  []QuizQuestion `json:"quiz_questions"`
}

type AnswerOptions struct {
	ID             int    `json:"id"`
	Value          string `json:"value"`
	Correct        bool   `json:"correct"`
	QuizQuestionID int    `json:"quiz_question_id"`
}

type QuizQuestion struct {
	ID            int             `json:"id"`
	QuizID        int             `json:"quiz_id"`
	Quiz          *Quiz           `json:"quiz"`
	QuestionText  string          `json:"question_text"`
	AnswerType    string          `json:"answer_type"`
	AnswerText    NullString      `json:"answer_text"`
	AnswerNumber  NullInt64       `json:"answer_number"`
	AnswerOptions []AnswerOptions `json:"answer_options"`
}

type PlayerAnswer struct {
	ID              int    `json:"id"`
	Answer          string `json:"answer"`
	PlayerID        Player `json:"player_id"`
	Timestamp       string `json:"timestamp"`
	TimestampClient string `json:"timestamp_client"`
}

type Tournament struct {
	ID             int    `json:"id"`
	TournamentName string `json:"tournament_name"`
}

type NullInt64 struct{ sql.NullInt64 }
type NullBool struct{ sql.NullBool }
type NullFloat64 struct{ sql.NullFloat64 }
type NullString struct{ sql.NullString }
type NullTime struct{ sql.NullTime }

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return nil
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

func (nb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return nil
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return nil
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return nil
}

// MarshalJSON for NullString
func (ns *NullTime) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Time)
}

// UnmarshalJSON for NullString
func (ns *NullTime) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.Time)
	if err != nil {
		fmt.Println("IMA GRESKA")
	}
	ns.Valid = (err == nil)
	return nil
}
