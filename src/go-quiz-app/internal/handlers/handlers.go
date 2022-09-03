package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/models"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/repository"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/repository/dbrepo"
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/utils"
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

func (m *Repository) GetQuiz(w http.ResponseWriter, r *http.Request) {
	quizCode := chi.URLParam(r, "code")

	quiz, err := m.DB.GetQuizInfoByCode(quizCode)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	fmt.Println(quiz)
	utils.ResponseJson(w, quiz)

}

func (m *Repository) GetQuizes(w http.ResponseWriter, r *http.Request) {

	var quizzes []models.Quiz
	var err error
	quizzes, err = m.DB.GetAllQuizes()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	// fmt.Printf("%+v \n", quizzes)
	// fmt.Printf("%+v \n", []models.Quiz{})
	if len(quizzes) == 0 {
		quizzes = []models.Quiz{}
	}

	for _, q := range quizzes {
		fmt.Println(q)
	}

	utils.ResponseJson(w, quizzes)

}

func (m *Repository) CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz
	json.NewDecoder(r.Body).Decode(&quiz)

	u := utils.CheckToken(r.Header.Get("Authorization"))
	quiz.OrganizerId = u.ID

	rand.Seed(time.Now().Unix())
	quiz.RoomCode = strconv.Itoa(rand.Intn(999999))

	newQuiz, err := m.DB.CreateQuiz(quiz)
	if err != nil {
		fmt.Println("ERROR CreateQuiz handler")
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	utils.ResponseJson(w, newQuiz)

}
