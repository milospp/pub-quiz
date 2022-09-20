package socketroom

import (
	"net"

	"github.com/milospp/pub-quiz/src/go-global/models"
)

var quiz_rooms map[int]*Room = make(map[int]*Room)

type ConnectedUser struct {
	Conn     net.Conn      `json:"-"`
	PlayerID int           `json:"player_id"`
	Player   models.Player `json:"player"`
	Room     *Room         `json:"-"`
}

type Room struct {
	QuizId   int
	Users    []*ConnectedUser `json:"users"`
	UsersMap map[string]*ConnectedUser
	Quiz     *models.Quiz
}

func GetRoom(id int) *Room {
	if val, ok := quiz_rooms[id]; ok {
		return val
	}

	return nil
}

func AddRoom(r *Room) {
	quiz_rooms[r.QuizId] = r
}

func GetPlayer(roomID int, conn *net.Conn) *models.Player {
	rm := GetRoom(roomID)
	if rm == nil {
		return nil
	}

	for _, u := range rm.Users {
		if u.Conn == *conn {
			return &u.Player
		}
	}
	return nil
}
