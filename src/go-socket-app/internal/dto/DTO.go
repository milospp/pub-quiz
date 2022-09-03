package dto

type Object map[string]interface{}

type SocketRequest struct {
	Username string `json:"username"`
	RoomID   int    `json:"room_id"`
	Method   string `json:"method"`
	Data     Object `json:"data"`
}

type SocketResponse struct {
	RoomID int    `json:"room_id"`
	Method string `json:"method"`
	Data   Object `json:"data"`
}
