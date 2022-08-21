package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/milospp/pub-quiz/pkg/config"
	"github.com/milospp/pub-quiz/pkg/driver"
	"github.com/milospp/pub-quiz/pkg/dto"
	"github.com/milospp/pub-quiz/pkg/repository"
	"github.com/milospp/pub-quiz/pkg/repository/dbrepo"
	"github.com/milospp/pub-quiz/pkg/utils"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("Connection IP: " + remoteIP)

	w.Write([]byte("Works!"))
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var logData dto.LoginDTO
	json.NewDecoder(r.Body).Decode(&logData)

	fmt.Println(logData)

	user, err := m.DB.Login(logData)
	if err != nil {
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":      user.Username,
		"email":         user.Email,
		"firstname":     user.FirstName,
		"lastname":      user.LastName,
		"team":          user.Team,
		"token_created": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("DontHeckMe"))
	utils.ResponseJson(w, tokenString)
}

func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	var regData dto.RegisterDTO
	json.NewDecoder(r.Body).Decode(&regData)

	var err error
	regData.Password, err = utils.HashPassword(regData.Password)
	if err != nil {
		return
	}

	fmt.Println("Reg: ")
	fmt.Println(regData)

	user, err := m.DB.Register(regData)
	if err != nil {
		log.Println(err)
		return
	}

	utils.ResponseJson(w, user)

}

func (m *Repository) GetQuizes(w http.ResponseWriter, r *http.Request) {

	quizzes, err := m.DB.GetAllQuizes()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Got quizes")
	fmt.Println(quizzes)

	for _, q := range quizzes {
		fmt.Println(q)
	}

	utils.ResponseJson(w, quizzes)

}
