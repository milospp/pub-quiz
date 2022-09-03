package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/milospp/pub-quiz/src/go-quiz-app/internal/models"
)

func (m *postgresDBRepo) GetQuizInfoByCode(code string) (models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, quiz_name, start_schedule, start_timestamp, end_timestamp, room_code, room_password, organizer_id FROM quizzes WHERE room_code=$1 LIMIT 1`
	var quiz models.Quiz

	row := m.DB.QueryRowContext(ctx, query, code)
	err := row.Scan(
		&quiz.ID,
		&quiz.QuizName,
		&quiz.StartSchedule,
		&quiz.StartTimestamp,
		&quiz.EndTimestamp,
		&quiz.RoomCode,
		&quiz.RoomPassword,
		&quiz.OrganizerId,
	)
	if err != nil {
		return quiz, err
	}

	return quiz, nil
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

func (m *postgresDBRepo) GetQuizById(id int) (models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var quiz models.Quiz

	query := `
		select q.id, q.quiz_name, q.start_schedule, q.start_timestamp, q.end_timestamp, q.room_code, q.room_password, q.organizer_id from quizzes q WHERE q.id = $1
	`

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
	)

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}

func (m *postgresDBRepo) GetQuestionByQuiz(id int) (models.QuizQuestion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var qq models.QuizQuestion

	query := `
		select q.id, q.quiz_id, q.question_text, q.answer_type, q.answer_text, q.answer_number from quizzes q WHERE q.quiz_id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&qq.ID,
		&qq.QuizID,
		&qq.QuestionText,
		&qq.AnswerType,
		&qq.AnswerText,
		&qq.AnswerNumber,
	)

	if err != nil {
		return qq, err
	}

	return qq, nil
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

func (m *postgresDBRepo) CreateQuiz(q models.Quiz) (models.Quiz, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	stmt := `INSERT INTO quizzes (quiz_name, organizer_id, start_schedule, start_timestamp, end_timestamp, room_code, room_password, created_at, updated_at)
			 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`

	row := m.DB.QueryRowContext(ctx, stmt,
		q.QuizName,
		q.OrganizerId,
		// q.StartSchedule,
		q.StartSchedule.Time,
		nil,
		nil,
		q.RoomCode,
		sql.NullString{},
		time.Now(),
		time.Now(),
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("Error CreateQuiz")
		return models.Quiz{}, err
	}

	q.ID = int(id)
	fmt.Println(id)
	questions := m.insertQuestion(q.QuizQuestions, q.ID, ctx)
	fmt.Println(questions)

	q.QuizQuestions = questions

	return q, nil

}

func (m *postgresDBRepo) insertQuestion(qs []models.QuizQuestion, qID int, ctx context.Context) []models.QuizQuestion {

	var questions []models.QuizQuestion

	for _, q := range qs {
		fmt.Println("ADDING QUIZ Q")
		fmt.Println(q)
		stmt := `INSERT INTO quiz_questions (quiz_id, question_text, answer_type, answer_text, answer_number, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

		row := m.DB.QueryRowContext(ctx, stmt,
			qID,
			q.QuestionText,
			q.AnswerType,
			q.AnswerText,
			q.AnswerNumber,
			time.Now(),
			time.Now(),
		)

		var id int
		err := row.Scan(&id)
		if err != nil {
			fmt.Println(err)
			fmt.Errorf("Cannot save question in quiz \n")
			continue
		}

		q.ID = id
		fmt.Println(q.ID)

		var answerOptions []models.AnswerOption
		for _, ao := range q.AnswerOpttions {
			ao = m.insertAnswerOption(ao, int(id), ctx)
			answerOptions = append(answerOptions, ao)
		}

		q.AnswerOpttions = answerOptions

		questions = append(questions, q)

	}

	return questions

}

func (m *postgresDBRepo) insertAnswerOption(ao models.AnswerOption, qqID int, ctx context.Context) models.AnswerOption {
	fmt.Println("ADDING Answer Option")

	stmt := `INSERT INTO answer_options (value, correct, quiz_question_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	row := m.DB.QueryRowContext(ctx, stmt,
		ao.Value,
		ao.Correct,
		qqID,
		time.Now(),
		time.Now(),
	)

	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Cannot save answer option")
		return models.AnswerOption{}
	}

	ao.ID = int(id)
	return ao
}
