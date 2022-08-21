package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/milospp/pub-quiz/pkg/dto"
	"github.com/milospp/pub-quiz/pkg/models"
	"github.com/milospp/pub-quiz/pkg/utils"
)

func (m *postgresDBRepo) Register(r dto.RegisterDTO) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (username, password, email, firstname, lastname, team, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.Username,
		r.Password,
		r.Email,
		r.FirstName,
		r.LastName,
		r.Team,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return models.User{}, err
	}

	l := dto.LoginDTO{
		Username: r.Username,
		Password: r.Password,
	}
	return m.Login(l)

}

func (m *postgresDBRepo) Login(l dto.LoginDTO) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, username, password, email, firstname, lastname, team FROM users WHERE username=$1 LIMIT 1`
	var user models.User

	row := m.DB.QueryRowContext(ctx, query, l.Username)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Team,
	)
	if err != nil {
		return user, err
	}

	if !utils.CheckPasswordHash(l.Password, user.Password) {
		return models.User{}, errors.New("Username or password not valid")
	}

	return user, nil
}

func (m *postgresDBRepo) GetAllQuizes() ([]models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var quizzes []models.Quiz

	query := `
		select q.id, q.quiz_name, q.start_schedule, q.start_timestamp, q.end_timestamp, q.room_code, q.room_password from quizzes q
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return quizzes, err
	}

	for rows.Next() {
		var quiz models.Quiz
		err := rows.Scan(
			&quiz.ID,
			&quiz.QuizName,
			&quiz.StartSchedule,
			&quiz.StartTimestamp,
			&quiz.EndTimestamp,
			&quiz.RoomCode,
			&quiz.RoomPassword,
		)
		if err != nil {
			return quizzes, err
		}
		quizzes = append(quizzes, quiz)
	}

	if err = rows.Err(); err != nil {
		return quizzes, err
	}

	return quizzes, nil
}

func (m *postgresDBRepo) GetScheduledQuizes() ([]models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var quizzes []models.Quiz

	query := `
		select q.id, q.quiz_name, q.start_schedule, q.start_timestamp, q.end_timestamp, q.room_code, q.room_password from quizzes q
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return quizzes, err
	}

	for rows.Next() {
		var quiz models.Quiz
		err := rows.Scan(
			&quiz.ID,
			&quiz.QuizName,
			&quiz.StartSchedule,
			&quiz.StartTimestamp,
			&quiz.EndTimestamp,
			&quiz.RoomCode,
			&quiz.RoomPassword,
			&quiz.QuizQuestions,
		)
		if err != nil {
			return quizzes, err
		}
		quizzes = append(quizzes, quiz)
	}

	if err = rows.Err(); err != nil {
		return quizzes, err
	}

	return quizzes, nil
}

func (*postgresDBRepo) GetMyQuizes() ([]models.Quiz, error) {
	return nil, nil

}

func (*postgresDBRepo) CreateQuiz() (models.Quiz, error) {
	return models.Quiz{}, nil
}
