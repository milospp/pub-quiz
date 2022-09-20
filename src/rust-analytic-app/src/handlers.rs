use std::collections::HashMap;

use postgres::{Client, Error, NoTls};
use serde::{Deserialize, Serialize};

use crate::{POSTGRES_DB, POSTGRES_HOST, POSTGRES_PASSWORD, POSTGRES_USER};

#[derive(Serialize, Deserialize)]
pub struct Stats {
    player_id: i32,
    correct_answers: i32,
    first_answers: i32,
    first_correct_answers: i32,
    total_answers: i32,
    points: i32,
}

#[derive(Serialize, Deserialize)]
pub struct AnswerOptions {
    id: i32,
    value: String,
    correct: bool,
    quiz_question_id: i32,
}

#[derive(Serialize, Deserialize)]
pub struct QuizQuestion {
    id: i32,
    quiz_id: i32,
    quiz: Option<Quiz>,
    question_text: String,
    answer_type: String,
    answer_text: Option<String>,
    answer_number: Option<i32>,
    answer_options: Vec<AnswerOptions>,
    player_answers: Vec<PlayerAnswer>,
}

#[derive(Serialize, Deserialize)]
pub struct Quiz {
    id: i32,
    quiz_name: String,
    start_schedule: Option<std::time::SystemTime>,
    start_timestamp: Option<std::time::SystemTime>,
    end_timestamp: Option<std::time::SystemTime>,
    room_code: String,
    room_password: Option<String>,
    organizer_id: Option<i32>,
    quiz_state: String,
    quiz_question: i32,
    question_state: i32,
    quiz_questions: Vec<QuizQuestion>,
    only_invited: bool,
}

#[derive(Serialize, Deserialize)]
pub struct PlayerAnswer {
    id: i32,
    answer: String,
    player_id: i32,
    question: Option<QuizQuestion>,
    question_id: i32,
    timestamp: Option<std::time::SystemTime>,
    timestamp_client: Option<std::time::SystemTime>,
    correct: bool,
    order: i32,
    correct_order: i32,
}

pub fn create_database_connection() -> Client {
    return Client::connect(
        format!(
            "postgresql://{}:{}@{}:5432/{}",
            POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB
        )
        .as_str(),
        NoTls,
    )
    .unwrap();
}

pub fn test_db() -> Result<String, Error> {
    let mut client = create_database_connection();

    for row in client.query("SELECT id FROM players", &[])? {
        let id: i32 = row.get(0);

        println!("found person: {}", id);
    }
    client.close()?;

    println!("Test");

    Ok("{test: 1}".to_string())
}

pub fn get_stats_quiz(quiz_id: i32) -> Result<String, Error> {
    let mut players: HashMap<i32, Stats> = HashMap::new();

    let questions = get_questions(quiz_id);
    for question in questions {
        for ans in question.player_answers.iter() {
            let mut stat = get_stats_answer(ans);
            if players.contains_key(&(stat.player_id)) {
                let hash_data = (players.get(&(stat.player_id)).unwrap().clone());
                stat.points += hash_data.points;
                stat.first_answers += hash_data.first_answers;
                stat.first_correct_answers += hash_data.first_correct_answers;
                stat.total_answers += hash_data.total_answers;
                stat.correct_answers += hash_data.correct_answers;

                players.insert(stat.player_id, stat);
            } else {
                players.insert(stat.player_id, stat);
            }
        }
    }

    Ok(serde_json::to_string(&players).unwrap())
}

// pub fn get_status_question(question: &QuizQuestion) {
//     for answer in question.player_answers {
//         if
//     }
// }

pub fn get_stats_answer(answer: &PlayerAnswer) -> Stats {
    let points;

    // TODO: Set min
    if answer.correct {
        points = 100 - (answer.correct_order * 5);
    } else {
        points = 0;
    }

    let first_ans;
    if answer.order == 0 {
        first_ans = 1;
    } else {
        first_ans = 0;
    }

    let first_correct_ans;
    if answer.correct && answer.correct_order == 0 {
        first_correct_ans = 1;
    } else {
        first_correct_ans = 0;
    }

    let correct;
    if answer.correct {
        correct = 1
    } else {
        correct = 0
    }

    Stats {
        player_id: answer.player_id,
        correct_answers: correct,
        first_answers: first_ans,
        first_correct_answers: first_correct_ans,
        total_answers: 1,
        points: points,
    }
}

pub fn get_full_quiz(quiz_id: i32) -> Result<String, Error> {
    let mut client = create_database_connection();

    let row = client.query_one("SELECT id, quiz_name, start_schedule, start_timestamp, end_timestamp, room_code, room_password, organizer_id,
    quiz_state,
    quiz_question,
    question_state FROM quizzes WHERE id=$1 LIMIT 1", &[&quiz_id]).unwrap();

    let quiz: Quiz = Quiz {
        id: row.get(0),
        quiz_name: row.get(1),
        start_schedule: row.get(2),
        start_timestamp: row.get(3),
        end_timestamp: row.get(4),
        room_code: row.get(5),
        room_password: row.get(6),
        organizer_id: row.get(7),
        quiz_state: row.get(8),
        quiz_question: row.get(9),
        question_state: row.get(10),
        only_invited: false,
        quiz_questions: get_questions(quiz_id),
    };

    Ok(serde_json::to_string(&quiz).unwrap())
}

pub fn get_questions(quiz_id: i32) -> Vec<QuizQuestion> {
    let mut client = create_database_connection();

    let mut questions: Vec<QuizQuestion> = Vec::new();

    let rows = client
        .query(
            "SELECT
            id,
            quiz_id,
            question_text,
            answer_type,
            answer_text,
            answer_number
            FROM quiz_questions WHERE quiz_id=$1",
            &[&quiz_id],
        )
        .unwrap();

    for row in rows {
        let mut quiz_question: QuizQuestion = QuizQuestion {
            id: row.get(0),
            quiz: Option::None,
            quiz_id: row.get(1),
            question_text: row.get(2),
            answer_type: row.get(3),
            answer_text: row.get(4),
            answer_number: row.get(5),
            answer_options: get_answer_options(row.get(0)),
            player_answers: get_player_answers(row.get(0)),
        };

        quiz_question
            .player_answers
            .sort_by(|a, b| a.timestamp.unwrap().cmp(&b.timestamp.unwrap()));
        validate_answer(&mut quiz_question);
        questions.push(quiz_question);
    }

    return questions;
}

pub fn get_answer_options(question_id: i32) -> Vec<AnswerOptions> {
    let mut client = create_database_connection();

    let mut answer_options: Vec<AnswerOptions> = Vec::new();

    let rows = client
        .query(
            "SELECT
            id,
            value,
            correct,
            quiz_question_id
            FROM answer_options WHERE quiz_question_id=$1",
            &[&question_id],
        )
        .unwrap();

    for row in rows {
        let answer_option: AnswerOptions = AnswerOptions {
            id: row.get(0),
            value: row.get(1),
            correct: row.get(2),
            quiz_question_id: row.get(3),
        };
        answer_options.push(answer_option);
    }

    return answer_options;
}

pub fn get_player_answers(question_id: i32) -> Vec<PlayerAnswer> {
    let mut client = create_database_connection();

    let mut player_answers: Vec<PlayerAnswer> = Vec::new();

    let rows = client
        .query(
            "SELECT
            id,
            answer,
            player_id,
            question_id,
            timestamp,
            timestamp_client
            FROM player_answer WHERE question_id=$1",
            &[&question_id],
        )
        .unwrap();

    for row in rows {
        let player_answer: PlayerAnswer = PlayerAnswer {
            id: row.get(0),
            answer: row.get(1),
            player_id: row.get(2),
            question: Option::None,
            question_id: row.get(3),
            timestamp: row.get(4),
            timestamp_client: row.get(5),
            correct: false,
            order: 0,
            correct_order: 0,
        };
        player_answers.push(player_answer);
    }

    return player_answers;
}

pub fn validate_answer(question: &mut QuizQuestion) {
    let correct_ans = get_correct_answer(question);

    let mut correct_index = 0;
    for (i, ans) in (*question).player_answers.iter_mut().enumerate() {
        if (*ans).answer == correct_ans {
            ans.correct = true;
            ans.correct_order = correct_index;
            correct_index += 1;
        }
        ans.order = i as i32;
    }
}

pub fn get_correct_answer(question: &QuizQuestion) -> String {
    match (*question).answer_type.as_str() {
        "SELECT" => {
            for (i, opt) in question.answer_options.iter().enumerate() {
                if opt.correct {
                    return i.to_string();
                }
            }

            return "".to_string();
        }
        "TEXT" => {
            let answer_text = question.answer_text.clone();
            answer_text.unwrap()
        }
        "NUMBER" => question.answer_number.unwrap().to_string(),
        _ => "".to_string(),
    }
}
