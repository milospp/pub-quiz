package repository

import (
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/dto"
	"github.com/milospp/pub-quiz/src/go-global/models"
)

type DatabaseRepo interface {
	Register(dto.RegisterDTO) (models.User, error)
	Login(dto.LoginDTO) (models.User, error)

	GetUserByID(int) (models.User, error)
	GetAnonymousUserByID(int) (models.AnonymousUser, error)
	RegisterAnonymous(models.AnonymousUser) (models.AnonymousUser, error)
}
