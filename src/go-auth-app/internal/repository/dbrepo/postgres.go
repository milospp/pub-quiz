package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/milospp/pub-quiz/src/go-auth-app/internal/dto"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/models"
	"github.com/milospp/pub-quiz/src/go-auth-app/internal/utils"
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

func (m *postgresDBRepo) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, username, password, email, firstname, lastname, team FROM users WHERE id=$1 LIMIT 1`
	var user models.User

	row := m.DB.QueryRowContext(ctx, query, id)
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

	return user, nil
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
