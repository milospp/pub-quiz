package repository

import (
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/models"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/utils"
)

type DatabaseRepo interface {
	// GetAllQuizes() ([]models.Quiz, error)
	// GetScheduledQuizes() ([]models.Quiz, error)
	// GetMyQuizes() ([]models.Quiz, error)
	// CreateQuiz() (models.Quiz, error)

	SetGameStates(utils.GameStateStruct) error
	GetQuestionById(int) (models.QuizQuestion, error)
	GetFullQuiz(int) (models.Quiz, error)
	GetQuestionsFull(int) ([]models.QuizQuestion, error)
	GetAnswerOptions(int) ([]models.AnswerOptions, error)

	// SetQuizGameState(int, string)
}
