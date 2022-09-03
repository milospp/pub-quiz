package socketroom

import (
	"fmt"
	"net"

	"github.com/milospp/pub-quiz/src/go-socket-app/internal/models"
)

var quiz_rooms map[int]*Room = make(map[int]*Room)

type User struct {
	// Io   sync.Mutex
	Conn net.Conn `json:"-"`

	UserId    int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Team      string `json:"team"`

	Room *Room `json:"-"`
}

type Room struct {
	QuizId int
	// Mu       sync.RWMutex
	Users    []*User `json:"users"`
	UsersMap map[string]*User
	Quiz     *models.Quiz
}

func GetRoom(id int) *Room {
	if val, ok := quiz_rooms[id]; ok {
		return val
	}

	r := &Room{
		QuizId: id,
	}

	quiz_rooms[id] = r
	return r
}

func GetUser(conn net.Conn) *User {
	fmt.Printf("%v", conn)
	fmt.Printf("%v", conn.LocalAddr().Network())
	return &User{}
}
