package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

const (
	GAME_PLAYER    int = 0
	GAME_SPECTATOR int = 1
	GAME_ADMIN     int = 2

	ROLE_USER  int = 0
	ROLE_ADMIN int = 1

	STATUS_OFFLINE     int = 0
	STATUS_ONLINE      int = 1
	STATUS_DISCONECTED int = 2
	STATUS_BANNED      int = 3

	QUIZ_TYPE_PRIVATE         int = 0
	QUIZ_TYPE_PUBLIC          int = 1
	QUIZ_TYPE_PUBLIC_APPROVED int = 21
	QUIZ_TYPE_PUBLIC_DECLINED int = 31
	QUIZ_TYPE_TOURNAMENT      int = 2
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Team      string `json:"team"`
	Role      int    `json:"role"`
}

type AnonymousUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Jwt  string `json:"jwt"`
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
	QuizType       int            `json:"quiz_type"`
	QuizState      string         `json:"quiz_state"`
	QuizQuestion   int            `json:"quiz_question"`
	QuestionState  int            `json:"question_state"`
	QuizQuestions  []QuizQuestion `json:"quiz_questions"`
	OnlyInvited    bool           `json:"only_invited"`
}

type Player struct {
	ID              int           `json:"id"`
	UserID          NullInt64     `json:"user_id"`
	User            User          `json:"user"`
	AnonymousUserID NullInt64     `json:"anonymous_user_id"`
	AnonymousUser   AnonymousUser `json:"anonymous_user"`
	Role            int           `json:"role"`
	Quiz            Quiz          `json:quiz`
	QuizID          int           `json:quiz_id`
	Team            string        `json:"team"`
	Status          int           `json:"status"`
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
	ID              int          `json:"id"`
	Answer          string       `json:"answer"`
	PlayerID        int          `json:"player_id"`
	Question        QuizQuestion `json:"question"`
	QuestionID      int          `json:"question_id"`
	Timestamp       string       `json:"timestamp"`
	TimestampClient string       `json:"timestamp_client"`
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
	fmt.Println(ni)
	fmt.Println("-----")

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

func (p *Player) GetFullName() string {
	if p.AnonymousUserID.Valid {
		return p.AnonymousUser.Name
	} else {
		return p.User.FirstName + " " + p.User.LastName
	}
}
