package repository

import (
	"github.com/milospp/pub-quiz/pkg/dto"
	"github.com/milospp/pub-quiz/pkg/models"
)

type DatabaseRepo interface {
	Register(dto.RegisterDTO) (models.User, error)
	Login(dto.LoginDTO) (models.User, error)

	GetAllQuizes() ([]models.Quiz, error)
	GetScheduledQuizes() ([]models.Quiz, error)
	GetMyQuizes() ([]models.Quiz, error)
	CreateQuiz() (models.Quiz, error)
}
