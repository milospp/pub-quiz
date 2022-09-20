package sockethandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/gobwas/ws/wsutil"
	"github.com/milospp/pub-quiz/src/go-global/models"
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

		r.registerUser(claim, conn, m.RoomID)
		r.broadcastLoggedUsers(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "START_GAME":
		// TODO: Validate
		r.setStartGame(m.RoomID)
		r.broadCastStartQuiz(m.RoomID)

		fmt.Println("START GAMEE")
		break

	case "SET_STATE":
		r.setStartGame(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "NEXT_STATE":
		r.nextQuestionState(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "RESTART":
		r.restartGame(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "FINISH":
		r.finishGame(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "SHOW_STATS":
		r.showStats(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "HIDE_STATS":
		r.hideStats(m.RoomID)
		r.broadcastQuizState(m.RoomID)
		break

	case "CHANGLE_PLAYER_ROLE":
		r.changePlayerRole(m.RoomID, m.Data)
		r.broadcastLoggedUsers(m.RoomID)
		break

	case "ANSWER":
		p := socketroom.GetPlayer(m.RoomID, &conn)
		r.updateAnswer(m.RoomID, *p, m.Data)

	case "DISCONNECTED":
		r.disconectedUser(m.RoomID, conn)
		r.broadcastLoggedUsers(m.RoomID)
		break
	}

}

func (r *Repository) GetOrCreateRoom(id int) *socketroom.Room {
	rm := socketroom.GetRoom(id)
	if rm != nil {
		return rm
	}

	rm = &socketroom.Room{
		QuizId: id,
	}

	quiz, err := r.DB.GetFullQuiz(id)
	if err != nil {
		fmt.Println(err)
	}

	rm.Quiz = &quiz

	socketroom.AddRoom(rm)
	return rm
}

func (r *Repository) nextQuestionState(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

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
	rm := r.GetOrCreateRoom(rmid)

	if rm.Quiz.QuizQuestion >= len(rm.Quiz.QuizQuestions)-1 {
		rm.Quiz.QuizState = "FINISHED"
	} else {
		rm.Quiz.QuizQuestion = rm.Quiz.QuizQuestion + 1
		rm.Quiz.QuestionState = 0
	}

	r.saveGameState(rm.Quiz)
	// fmt.Println("Setted state Quiz")

}

func (r *Repository) saveGameState(quiz *models.Quiz) {
	qs := utils.GameStateStruct{
		QuizID:        quiz.ID,
		QuizState:     quiz.QuizState,
		QuizQuestion:  quiz.QuizQuestion,
		QuestionState: quiz.QuestionState,
	}

	err := r.DB.SetGameStates(qs)
	if err != nil {
		fmt.Println(err)
		// TODO: Broadcast error
		return
	}

}

func (r *Repository) setStartGame(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

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

func (r *Repository) restartGame(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

	rm.Quiz.QuizState = "LOBBY"
	rm.Quiz.QuizQuestion = 0
	rm.Quiz.QuestionState = 0
	r.saveGameState(rm.Quiz)

}

func (r *Repository) finishGame(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

	rm.Quiz.QuizState = "FINISHED"
	rm.Quiz.QuizQuestion = len(rm.Quiz.QuizQuestions) - 1
	rm.Quiz.QuestionState = 4

	r.saveGameState(rm.Quiz)
}

func (r *Repository) showStats(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

	rm.Quiz.QuizState = "STATS"

	r.saveGameState(rm.Quiz)
}

func (r *Repository) hideStats(rmid int) {
	rm := r.GetOrCreateRoom(rmid)

	if rm.Quiz.QuizQuestion == len(rm.Quiz.QuizQuestions)-1 && rm.Quiz.QuestionState == 4 {
		rm.Quiz.QuizState = "FINISHED"
	} else {
		rm.Quiz.QuizState = "QUIZ"
	}

	r.saveGameState(rm.Quiz)
}

func AuthUser(m dto.SocketRequest, conn net.Conn) *utils.AuthClaim {
	jwt := m.Data["jwt"].(string)
	claim := utils.CheckToken(jwt)

	return claim
}

func (r *Repository) broadCastStartQuiz(roomId int) {
	res := dto.SocketResponse{
		RoomID: roomId,
		Method: "START_GAME",
	}

	data, _ := json.Marshal(res)

	r.broadcastEveryione(roomId, data)

}

func (r *Repository) broadcastQuizState(roomId int) {
	rm := r.GetOrCreateRoom(roomId)

	d := make(dto.Object)
	d["game_state"] = utils.GameStateStruct{
		QuizID:        rm.Quiz.ID,
		QuizState:     rm.Quiz.QuizState,
		QuizQuestion:  rm.Quiz.QuizQuestion,
		QuestionState: rm.Quiz.QuestionState,
	}

	if rm.Quiz.QuizQuestion < len(rm.Quiz.QuizQuestions) {
		d["question_id"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion].ID

		if rm.Quiz.QuestionState == 1 {
			d["question_text"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion].QuestionText
			d["answer_type"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion].AnswerType
		}

		if rm.Quiz.QuestionState == 2 {
			d["question"] = rm.Quiz.QuizQuestions[rm.Quiz.QuizQuestion]
		}
	}

	d["game_state"] = rm.Quiz

	res := dto.SocketResponse{
		RoomID: roomId,
		Method: "QUIZ_STATE",
		Data:   d,
	}

	data, _ := json.Marshal(res)

	r.broadcastEveryione(roomId, data)
}

func (r *Repository) broadcastEveryione(roomId int, data []byte) {
	rm := r.GetOrCreateRoom(roomId)
	for _, v := range rm.Users {
		wsutil.WriteServerMessage(v.Conn, 1, data)
	}
}

func (r *Repository) registerUser(claim *utils.AuthClaim, conn net.Conn, roomId int) *socketroom.ConnectedUser {
	room := r.GetOrCreateRoom(roomId)

	oldSu := checkIfPlayerConnected(claim, room)
	if oldSu.PlayerID > 0 {
		// AMEND CONNECTION
		oldSu.Conn = conn
		oldSu.Player.Status = models.STATUS_ONLINE
		r.DB.UpdatePlayer(oldSu.Player)
		return oldSu
	}

	su := socketroom.ConnectedUser{
		Conn: conn,
		Room: room,
	}

	pl := models.Player{
		Role:   models.GAME_PLAYER,
		QuizID: room.QuizId,
		Status: models.STATUS_ONLINE,
	}

	if claim.AnonymousKey == "" && room.Quiz.OrganizerId == claim.ID {
		pl.Role = models.GAME_ADMIN
	}

	var nulInt models.NullInt64
	nulInt.Int64 = int64(claim.ID)
	nulInt.Valid = true

	if claim.AnonymousKey != "" {
		pl.AnonymousUserID = nulInt
		pl.AnonymousUser = models.AnonymousUser{
			ID:   claim.ID,
			Name: claim.FirstName,
		}
	} else {
		pl.UserID = nulInt
		pl.User = models.User{
			ID:        claim.ID,
			FirstName: claim.FirstName,
			LastName:  claim.Lastname,
			Role:      claim.Role,
		}
	}

	pl, err := r.addPlayerToDB(pl)

	// TODO: sta ako je greska
	if err != nil {
		return &su
	}

	su.PlayerID = pl.ID
	su.Player = pl
	room.Users = append(room.Users, &su)

	return &su

}

func checkIfPlayerConnected(claim *utils.AuthClaim, room *socketroom.Room) *socketroom.ConnectedUser {
	isAnon := false
	if claim.AnonymousKey != "" {
		isAnon = true
	}

	var id = claim.ID

	for _, v := range room.Users {
		if isAnon && v.Player.AnonymousUserID.Valid && int(v.Player.AnonymousUserID.Int64) == id {
			return v
		} else if int(v.Player.UserID.Int64) == id {
			return v
		}

	}
	return &socketroom.ConnectedUser{}

}

func (r *Repository) addPlayerToDB(pl models.Player) (models.Player, error) {
	pl, err := r.DB.InsertPlayer(pl)
	return pl, err
}

func (r *Repository) broadcastLoggedUsers(roomId int) {
	rm := r.GetOrCreateRoom(roomId)
	fmt.Println("Broadcasting all users from room")

	data := getUserListInRoom(rm)
	for _, v := range rm.Users {
		wsutil.WriteServerMessage(v.Conn, 1, data)
	}
}

func broadcastQuestion() {

}

func getUserListInRoom(room *socketroom.Room) []byte {
	d := make(dto.Object)

	playersList := []models.Player{}
	for _, v := range room.Users {
		playersList = append(playersList, v.Player)
	}
	d["players"] = playersList

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

func (r *Repository) disconectedUser(rmid int, conn net.Conn) (models.Player, error) {
	rm := r.GetOrCreateRoom(rmid)

	for _, v := range rm.Users {
		if v.Conn == conn {
			fmt.Printf("User %d disconected\n", v.Player.ID)
			v.Player.Status = models.STATUS_OFFLINE
			r.DB.UpdatePlayer(v.Player)
			return v.Player, nil

		}
	}

	return models.Player{}, errors.New("Not found user")

}

func (r *Repository) changePlayerRole(rmid int, data dto.Object) (models.Player, error) {
	rm := r.GetOrCreateRoom(rmid)

	for _, v := range rm.Users {
		if v.PlayerID == int(data["player_id"].(float64)) {
			v.Player.Role = int(data["role"].(float64))
			r.DB.UpdatePlayer(v.Player)
			return v.Player, nil

		}
	}

	return models.Player{}, errors.New("Not found user")

}

func (r *Repository) updateAnswer(rmid int, player models.Player, data dto.Object) {
	// rm := r.GetOrCreateRoom(rmid)

	fmt.Println("Updating answer")
	answer := models.PlayerAnswer{
		Answer:          data["answer"].(string),
		PlayerID:        player.ID,
		Timestamp:       time.Now().Format("2006-01-02T15:04:05.999Z07:00"),
		TimestampClient: data["timestamp"].(string),
		QuestionID:      int(data["question_id"].(float64)),
	}
	err := r.DB.UpdatePlayerAnswer(answer)
	if err != nil {
		answer, err = r.DB.InsertPlayerAnswer(answer)
	}
	if err != nil {
		fmt.Println("error updating answer")
	}

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
