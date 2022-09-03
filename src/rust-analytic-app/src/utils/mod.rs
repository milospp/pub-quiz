use chrono::{prelude::*, Duration};
use postgres::{Client, Error, NoTls};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Quiz {
    id: i32,
    quiz_name: String,
    start_schedule: String,
    start_timestamp: String,
    end_timestamp: String,
    room_code: String,
    room_password: String,
    organizer_id: i32,
    quiz_state: String,
    quiz_question: i32,
    question_state: i32,
}

// #[derive(Serialize, Deserialize)]
// pub struct QuizQuestion {
//     id: i32
//     quiz_id: i32
//     // quiz: i32
//     question_text: String
//     answer_type: String
//     answer_text: String
//     answer_number: i32
//     // answer_options: i32
// }

// #[derive(Serialize, Deserialize)]
// pub struct QuizQuestion {
//     id: i32
//     value: i32
//     correct: bool
//     quiz_question_id: String
// }

// #[derive(Serialize, Deserialize)]
// pub struct PlayerAnswer {
//     id: i32,
//     answer: String,
//     player_id: i32,
//     timestamp: String,
//     timestamp_client: String,
// }

#[derive(Serialize)]
struct Response {
    message: String,
}

pub fn get_qll_quizzes(id: i32) -> Result<(String), Error> {
    println!("Test get all quizzes");

    let mut client = Client::connect("postgresql://postgres:root@localhost:5432/pubquiz", NoTls)?;
    let mut ret: Vec<Quiz> = vec![];
    for row in client.query("SELECT id, quiz_name, start_schedule, start_timestamp, end_timestamp, room_code, room_password, organizer_id  FROM reservations WHERE roomId = $1 and cancelled = false", &[&id])? {
    let id: i32 = row.get(0);
    let quiz_name: String = row.get(1);
    let start_schedule: String = row.get(2);
    let start_timestamp: String = row.get(3);
    let end_timestamp: String = row.get(4);
    let room_code: String = row.get(5);
    let room_password: String = row.get(6);
    let organizer_id: i32 = row.get(8);
    let quiz_state: String = row.get(9);
    let quiz_question: i32 = row.get(10);
    let question_state: i32 = row.get(11);

    ret.push(Quiz { id: (id), quiz_name: (quiz_name), start_schedule: (start_schedule), start_timestamp: (start_timestamp), end_timestamp: (end_timestamp), room_code: (room_code), room_password: (room_password), organizer_id: (organizer_id), quiz_state: (quiz_state), quiz_question: (quiz_question), question_state: (question_state)});


}
    client.close()?;

    Ok(serde_json::to_string(&ret).unwrap())
}

// let mut ret: Vec<Quiz> = vec![];
// for row in client.query("SELECT quiz_name, start_schedule, start_timestamp, end_timestamp, room_code, room_password, organizer_id FROM quizzes", &[&id])? {
//     let quiz_name: i32 = row.get(0);
//     let quiz_name: &str = row.get(1);
//     let start_schedule: &str = row.get(2);
//     let start_timestamp: &str = row.get(3);
//     let end_timestamp: &str = row.get(4);
//     let room_code: &str = row.get(5);
//     let room_password: &str = row.get(6);
//     let organizer_id: &str = row.get(7);
//     ret.push(Quiz { id: (id), user_id: (user_id), room_id: (room_id), cancelled: (cancelled), start: (start.to_string()), end: (end.to_string()), username: (username.to_string()), hotel_name: (hotel_name.to_string()), room_number: (room_number) });
// }
// client.close()?;

// Ok("okej")
// Ok(serde_json::to_string(&ret).unwrap())
// }
