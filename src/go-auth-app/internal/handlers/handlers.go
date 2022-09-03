package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/dto"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/repository"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/repository/dbrepo"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/utils"
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

	fmt.Println(logData)

	user, err := m.DB.Login(logData)
	if err != nil {
		w.WriteHeader(401)
		log.Println("Wrong password")
		log.Println(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"team":       user.Team,
		"issued_at":  time.Now(),
		"expired_at": time.Now().Add(time.Hour * 48),
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

func (m *Repository) GetLoggedUser(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.Header.Get("Authorization"), " ")[1]
	fmt.Println(a)

	// var customClaim dto.AuthClaim

	token, err := jwt.ParseWithClaims(a, &config.AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte("DontHeckMe"), nil
	})
	fmt.Println(err)

	claims := token.Claims.(*config.AuthClaim)

	user, err := m.DB.GetUserById(claims.ID)
	if err != nil {
		return
	}

	utils.ResponseJson(w, user)

}
