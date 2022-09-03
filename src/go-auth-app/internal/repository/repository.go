package repository

import (
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/dto"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/models"
)

type DatabaseRepo interface {
	Register(dto.RegisterDTO) (models.User, error)
	Login(dto.LoginDTO) (models.User, error)

	GetUserById(int) (models.User, error)
}
