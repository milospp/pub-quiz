package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaim struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstname`
	Lastname  string    `json:"lastname"`
	Team      string    `json:"team"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type GameStateStruct struct {
	QuizID        int    `json:"quiz_id"`
	QuizState     string `json:"quiz_state"`
	QuizQuestion  int    `json:"quiz_question"`
	QuestionState int    `json:"question_state"`
}

func (claim *AuthClaim) Valid() error {
	if time.Now().After(claim.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}

func ResponseJson(w http.ResponseWriter, d interface{}) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "    ")
	e.Encode(d)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckToken(rawToken string) *AuthClaim {
	tok := strings.Split(rawToken, " ")
	if len(tok) < 2 {
		return nil
	}
	a := tok[1]

	token, err := jwt.ParseWithClaims(a, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("DontHeckMe"), nil
	})
	if err != nil {
		return nil
	}
	if token.Valid == false {
		return nil
	}

	return token.Claims.(*AuthClaim)
}
