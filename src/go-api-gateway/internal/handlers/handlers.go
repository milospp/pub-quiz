package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/milospp/pub-quiz/src/go-api-gateway/internal/config"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) RedirectToAuth(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest(r.Method, "http://localhost:8003"+r.URL.String(), r.Body)
	req.Header = r.Header

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func (m *Repository) RedirectToQuiz(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest(r.Method, "http://localhost:8001"+r.URL.String(), r.Body)
	req.Header = r.Header

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func (m *Repository) RedirectToStats(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest(r.Method, "http://localhost:8004"+r.URL.String(), r.Body)
	req.Header = r.Header

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

// // func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

// // 	http.Redirect(w, r, "localhost:8003"+"/", http.StatusSeeOther)

// // }

// func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

// 	// response, err := http.Post("localhost:8001" + "/login")
// 	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8003"+"/login", r.Body)
// 	client := &http.Client{}
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusGatewayTimeout)
// 		return
// 	}

// 	w.WriteHeader(response.StatusCode)
// 	io.Copy(w, response.Body)
// 	response.Body.Close()

// }

// func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {

// 	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8003"+"/register", r.Body)
// 	client := &http.Client{}
// 	response, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusGatewayTimeout)
// 		return
// 	}

// 	w.WriteHeader(response.StatusCode)
// 	io.Copy(w, response.Body)
// 	response.Body.Close()

// }

// func (m *Repository) GetLoggedUser(w http.ResponseWriter, r *http.Request) {

// 	response, err := http.Get("http://localhost:8003" + "/profile")
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusGatewayTimeout)
// 		return
// 	}
// 	w.WriteHeader(response.StatusCode)
// 	io.Copy(w, response.Body)
// 	response.Body.Close()

// }
