package dbrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/milospp/pub-quiz/src/go-global/models"
	"github.com/milospp/pub-quiz/src/go-socket-app/internal/utils"
)

// func (m *postgresDBRepo) Register(r dto.RegisterDTO) (models.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	stmt := `INSERT INTO users (username, password, email, firstname, lastname, team, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)`

// 	_, err := m.DB.ExecContext(ctx, stmt,
// 		r.Username,
// 		r.Password,
// 		r.Email,
// 		r.FirstName,
// 		r.LastName,
// 		r.Team,
// 		time.Now(),
// 		time.Now(),
// 	)

// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	l := dto.LoginDTO{
// 		Username: r.Username,
// 		Password: r.Password,
// 	}
// 	return m.Login(l)

// }

// func (m *postgresDBRepo) Login(l dto.LoginDTO) (models.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := `SELECT id, username, password, email, firstname, lastname, team FROM users WHERE username=$1 LIMIT 1`
// 	var user models.User

// 	row := m.DB.QueryRowContext(ctx, query, l.Username)
// 	err := row.Scan(
// 		&user.ID,
// 		&user.Username,
// 		&user.Password,
// 		&user.Email,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.Team,
// 	)
// 	if err != nil {
// 		return user, err
// 	}

// 	if !utils.CheckPasswordHash(l.Password, user.Password) {
// 		return models.User{}, errors.New("Username or password not valid")
// 	}

// 	return user, nil
// }

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

func (m *postgresDBRepo) SetGameStates(structs utils.GameStateStruct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		UPDATE quizzes
		SET quiz_state = $1,
		 	quiz_question = $2,
		 	question_state = $3,
			updated_at = $4
		WHERE id = $5
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		structs.QuizState,
		structs.QuizQuestion,
		structs.QuestionState,
		time.Now(),
		structs.QuizID,
	)

	return err

}

func (m *postgresDBRepo) GetQuestionById(id int) (models.QuizQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, question_text, answer_type, answer_text, answer_number, quiz_id FROM quiz_question WHERE id=$1`
	var qq models.QuizQuestion

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&qq.ID,
		&qq.QuestionText,
		&qq.AnswerType,
		&qq.AnswerText,
		&qq.AnswerNumber,
		&qq.QuizID,
	)
	if err != nil {
		return qq, err
	}

	query = `SELECT id, value, correct, quiz_question_id FROM answer_options WHERE quiz_question_id=$1`

	var aos []models.AnswerOptions

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return qq, err
	}

	for rows.Next() {
		var ao models.AnswerOptions
		err := rows.Scan(
			&ao.ID,
			&ao.Value,
			&ao.Correct,
			&ao.QuizQuestionID,
		)
		if err != nil {
			return qq, err
		}
		aos = append(aos, ao)
	}

	qq.AnswerOptions = aos

	return qq, nil
}

func (m *postgresDBRepo) GetFullQuiz(id int) (models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, quiz_name, start_schedule, start_timestamp, end_timestamp, room_code, room_password, organizer_id, quiz_state, quiz_question, question_state 
				FROM quizzes WHERE id=$1 LIMIT 1`

	var quiz models.Quiz

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&quiz.ID,
		&quiz.QuizName,
		&quiz.StartSchedule,
		&quiz.StartTimestamp,
		&quiz.EndTimestamp,
		&quiz.RoomCode,
		&quiz.RoomPassword,
		&quiz.OrganizerId,
		&quiz.QuizState,
		&quiz.QuizQuestion,
		&quiz.QuestionState,
	)

	if err != nil {
		return quiz, err
	}
	quiz.QuizQuestions, err = m.GetQuestionsFull(id)

	return quiz, nil
}

func (m *postgresDBRepo) GetQuestionsFull(quizID int) ([]models.QuizQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var quizQuestions []models.QuizQuestion

	query := `
		SELECT q.id, q.quiz_id, q.question_text, q.answer_type, q.answer_text, q.answer_number from quiz_questions q WHERE q.quiz_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, quizID)
	if err != nil {
		return quizQuestions, err
	}

	for rows.Next() {
		var qq models.QuizQuestion
		err := rows.Scan(
			&qq.ID,
			&qq.QuizID,
			&qq.QuestionText,
			&qq.AnswerType,
			&qq.AnswerText,
			&qq.AnswerNumber,
		)
		if err != nil {
			return quizQuestions, err
		}

		qq.AnswerOptions, err = m.GetAnswerOptions(qq.ID)
		quizQuestions = append(quizQuestions, qq)
	}
	if err != nil {
		return quizQuestions, err
	}

	return quizQuestions, nil
}

func (m *postgresDBRepo) GetAnswerOptions(quizQuestionID int) ([]models.AnswerOptions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var answerOptions []models.AnswerOptions

	query := `
		SELECT id, value, correct, quiz_question_id FROM answer_options WHERE quiz_question_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, quizQuestionID)
	if err != nil {
		return answerOptions, err
	}

	for rows.Next() {
		var ao models.AnswerOptions
		err := rows.Scan(
			&ao.ID,
			&ao.Value,
			&ao.Correct,
			&ao.QuizQuestionID,
		)
		if err != nil {
			return answerOptions, err
		}

		answerOptions = append(answerOptions, ao)
	}
	if err != nil {
		return answerOptions, err
	}

	return answerOptions, nil
}

func (m *postgresDBRepo) InsertPlayer(p models.Player) (models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stmt := `INSERT INTO players (role, quiz_id, user_id, anonymous_user_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	row := m.DB.QueryRowContext(ctx, stmt,
		p.Role,
		p.QuizID,
		p.UserID,
		p.AnonymousUserID,
		p.Status,
		time.Now(),
		time.Now(),
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("Error insert player")
		fmt.Println(err)
		return models.Player{}, err
	}

	p.ID = int(id)
	return p, nil
}

func (m *postgresDBRepo) UpdatePlayer(p models.Player) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stmt := `
		UPDATE players
		SET role = $2,
		quiz_id = $3,
		user_id = $4,
		anonymous_user_id = $5
		status = $6
		updated_at = $7
		WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		p.ID,
		p.Role,
		p.QuizID,
		p.UserID,
		p.AnonymousUserID,
		p.Status,
		time.Now(),
	)
	fmt.Println(err)

	return err

}

//////////////////////

func (m *postgresDBRepo) GetQuestionPlayerAnswers(qID int) ([]models.PlayerAnswer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, answer, player_id, question_id, timestamp, timestamp_client FROM player_answer WHERE quiz_id=$1`

	var pas []models.PlayerAnswer

	rows, err := m.DB.QueryContext(ctx, query, qID)
	if err != nil {
		return pas, err
	}

	for rows.Next() {
		var pa models.PlayerAnswer
		err := rows.Scan(
			&pa.ID,
			&pa.Answer,
			&pa.PlayerID,
			&pa.QuestionID,
			&pa.Timestamp,
			&pa.TimestampClient,
		)
		if err != nil {
			return pas, err
		}
		pas = append(pas, pa)
	}

	return pas, nil
}

func (m *postgresDBRepo) InsertPlayerAnswer(pa models.PlayerAnswer) (models.PlayerAnswer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stmt := `INSERT INTO player_answer (answer, player_id, question_id, timestamp, timestamp_client, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	row := m.DB.QueryRowContext(ctx, stmt,
		pa.Answer,
		pa.PlayerID,
		pa.QuestionID,
		pa.Timestamp,
		pa.TimestampClient,
		time.Now(),
		time.Now(),
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("Error insert player answer")
		fmt.Println(err)
		return models.PlayerAnswer{}, err
	}

	pa.ID = int(id)
	return pa, nil
}

func (m *postgresDBRepo) UpdatePlayerAnswer(pa models.PlayerAnswer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stmt := `
		UPDATE player_answer
		SET answer = $2,
		player_id = $3,
		question_id = $4,
		timestamp = $5,
		timestamp_client = $6,
		updated_at = $1
		WHERE player_id = $3 AND question_id = $4
	`

	updated, err := m.DB.ExecContext(ctx, stmt,
		time.Now(),
		pa.Answer,
		pa.PlayerID,
		pa.QuestionID,
		pa.Timestamp,
		pa.TimestampClient,
	)
	if err != nil {
		fmt.Println(err)
	}
	affected, err := updated.RowsAffected()
	if err != nil || affected == 0 {
		return errors.New("")
	}
	return err

}
