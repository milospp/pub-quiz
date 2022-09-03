package sockethandlers

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gobwas/ws/wsutil"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/config"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/driver"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/dto"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/repository"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/repository/dbrepo"
	socketroom "github.com/milospp/pub-quiz/src/go-socket-app/internal/socket-room"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/utils"
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

func (r *Repository) ProcessMessage(m dto.SocketRequest, conn net.Conn) {
	// fmt.Printf("PRC: %v \n", m.RoomID)

	switch m.Method {
	case "AUTH":
		// socketroom.GetUser(conn)

		claim := AuthUser(m, conn)
		if claim == nil {
			return
		}

		registerUser(claim, conn, m.RoomID)
		broadcastLoggedUsers(m.RoomID)

		break
	case "START_GAME":
		// TODO: Validate
		r.setStartGame(m.RoomID)
		broadCastStartQuiz(m.RoomID)

		fmt.Println("START GAMEE")
		break

	case "SET_STATE":
		r.setStartGame(m.RoomID)
		broadcastQuizState(m.RoomID)
		break

	case "NEXT_STATE":
		r.nextQuestionState(m.RoomID)
		broadcastQuizState(m.RoomID)
		break

	case "ANSWER":
		break

	}

}

func (r *Repository) nextQuestionState(rmid int) {
	rm := socketroom.GetRoom(rmid)

	if rm.Quiz.QuestionState >= 4 {
		r.nextQuizQuestion(rmid)
	} else {
		rm.Quiz.QuestionState = rm.Quiz.QuestionState + 1
	}

	qs := utils.GameStateStruct{
		QuizID:        rm.Quiz.ID,
		QuizState:     rm.Quiz.QuizState,
		QuizQuestion:  rm.Quiz.QuizQuestion,
		QuestionState: rm.Quiz.QuestionState,
	}

	err := r.DB.SetGameStates(qs)
	if err != nil {
		fmt.Println(err)
		// TODO: Broadcast error
		return
	}

	// fmt.Println("Setted state Quiz")

}

func (r *Repository) nextQuizQuestion(rmid int) {
	rm := socketroom.GetRoom(rmid)

	rm.Quiz.QuizQuestion = rm.Quiz.QuizQuestion + 1
	rm.Quiz.QuestionState = 0

	qs := utils.GameStateStruct{
		QuizID:        rm.Quiz.ID,
		QuizState:     rm.Quiz.QuizState,
		QuizQuestion:  rm.Quiz.QuizQuestion,
		QuestionState: rm.Quiz.QuestionState,
	}

	err := r.DB.SetGameStates(qs)
	if err != nil {
		fmt.Println(err)
		// TODO: Broadcast error
		return
	}

	// fmt.Println("Setted state Quiz")

}

func (r *Repository) setStartGame(rmid int) {
	rm := socketroom.GetRoom(rmid)

	quiz, err := r.DB.GetFullQuiz(rmid)
	if err != nil {
		fmt.Println(err)
	}
	rm.Quiz = &quiz

	s := utils.GameStateStruct{
		QuizID:        rm.QuizId,
		QuizState:     "QUIZ",
		QuizQuestion:  0,
		QuestionState: 0,
	}

	fmt.Println(s)

	rm.Quiz.QuizState = s.QuizState
	rm.Quiz.QuizQuestion = s.QuizQuestion
	rm.Quiz.QuestionState = s.QuestionState

	err = r.DB.SetGameStates(s)
	if err != nil {
		fmt.Println(err)
		// TODO: Broadcast error
		return
	}

	// TODO: FIX sh*t

	fmt.Println("Setted state Quiz")

}

func AuthUser(m dto.SocketRequest, conn net.Conn) *utils.AuthClaim {
	jwt := m.Data["jwt"].(string)
	claim := utils.CheckToken(jwt)

	return claim
}

func broadCastStartQuiz(roomId int) {
	res := dto.SocketResponse{
		RoomID: roomId,
		Method: "START_GAME",
	}

	r, _ := json.Marshal(res)

	broadcastEveryione(roomId, r)

}

func broadcastQuizState(roomId int) {
	rm := socketroom.GetRoom(roomId)

	fmt.Println(rm.Quiz.QuizQuestion)
	fmt.Println(rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion])

	d := make(dto.Object)
	d["game_state"] = utils.GameStateStruct{
		QuizID:        rm.Quiz.ID,
		QuizState:     rm.Quiz.QuizState,
		QuizQuestion:  rm.Quiz.QuizQuestion,
		QuestionState: rm.Quiz.QuestionState,
	}

	if rm.Quiz.QuestionState == 1 {
		d["question_text"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion].QuestionText
		d["answer_type"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion].AnswerType
	}

	if rm.Quiz.QuestionState == 2 {
		d["question"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion]
	}

	d["game_state"] = rm.Quiz

	res := dto.SocketResponse{
		RoomID: roomId,
		Method: "QUIZ_STATE",
		Data:   d,
	}

	r, _ := json.Marshal(res)

	broadcastEveryione(roomId, r)
}

func broadcastEveryione(roomId int, data []byte) {
	r := socketroom.GetRoom(roomId)
	for _, v := range r.Users {
		wsutil.WriteServerMessage(v.Conn, 1, data)
	}
}

func registerUser(claim *utils.AuthClaim, conn net.Conn, roomId int) *socketroom.User {
	r := socketroom.GetRoom(roomId)

	su := socketroom.User{
		UserId:    claim.ID,
		Username:  claim.Username,
		FirstName: claim.FirstName,
		LastName:  claim.Lastname,
		Team:      claim.Team,
		Conn:      conn,
		Room:      r,
	}

	r.Users = append(r.Users, &su)

	return &su

}

func broadcastLoggedUsers(roomId int) {
	r := socketroom.GetRoom(roomId)
	fmt.Println("Broadcasting all users from room")
	for _, v := range r.Users {
		wsutil.WriteServerMessage(v.Conn, 1, getUserListInRoom(r))
	}
}

func broadcastQuestion() {

}

func getUserListInRoom(room *socketroom.Room) []byte {
	d := make(dto.Object)
	d["users"] = room.Users

	resObj := dto.SocketResponse{
		RoomID: room.QuizId,
		Method: "UPDATE_USERS",
		Data:   d,
	}

	r, err := json.Marshal(resObj)

	if err != nil {
		return []byte("[]")
	}

	return r

}

// func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
// 	remoteIP := r.RemoteAddr
// 	fmt.Println("Connection IP: " + remoteIP)

// 	w.Write([]byte("Works!"))
// }

// func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
// 	var logData dto.LoginDTO
// 	json.NewDecoder(r.Body).Decode(&logData)

// 	fmt.Println(logData)

// 	user, err := m.DB.Login(logData)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username":      user.Username,
// 		"email":         user.Email,
// 		"firstname":     user.FirstName,
// 		"lastname":      user.LastName,
// 		"team":          user.Team,
// 		"token_created": time.Now().Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte("DontHeckMe"))
// 	utils.ResponseJson(w, tokenString)
// }

// func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
// 	var regData dto.RegisterDTO
// 	json.NewDecoder(r.Body).Decode(&regData)

// 	var err error
// 	regData.Password, err = utils.HashPassword(regData.Password)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println("Reg: ")
// 	fmt.Println(regData)

// 	user, err := m.DB.Register(regData)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	utils.ResponseJson(w, user)

// }

// func (m *Repository) GetQuizes(w http.ResponseWriter, r *http.Request) {

// 	quizzes, err := m.DB.GetAllQuizes()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Got quizes")
// 	fmt.Println(quizzes)

// 	for _, q := range quizzes {
// 		fmt.Println(q)
// 	}

// 	utils.ResponseJson(w, quizzes)

// }
