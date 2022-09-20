package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var randLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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
func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randLetters[rand.Intn(len(randLetters))]
	}
	return string(b)
}
