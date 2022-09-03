package repository

import (
	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/models"
)

type DatabaseRepo interface {
	GetQuizInfoByCode(string) (models.Quiz, error)
	GetAllQuizes() ([]models.Quiz, error)
	GetScheduledQuizes() ([]models.Quiz, error)
	GetMyQuizes() ([]models.Quiz, error)
	CreateQuiz(models.Quiz) (models.Quiz, error)
}
