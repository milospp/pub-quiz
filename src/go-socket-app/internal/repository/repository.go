package repository

import (
	"github.com/milospp/pub-quiz/src/go-global/models"
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

	InsertPlayer(models.Player) (models.Player, error)
	UpdatePlayer(models.Player) error

	GetQuestionPlayerAnswers(int) ([]models.PlayerAnswer, error)
	InsertPlayerAnswer(models.PlayerAnswer) (models.PlayerAnswer, error)
	UpdatePlayerAnswer(models.PlayerAnswer) error

	// SetQuizGameState(int, string)
}
