package dto

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Team      string `json:"team"`
}
