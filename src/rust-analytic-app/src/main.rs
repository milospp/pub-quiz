use std::thread;

use rocket::response::content;
use rust_analytic_app::handlers;

#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> content::RawJson<String> {
    thread::spawn(move || content::RawJson(handlers::test_db().unwrap()))
        .join()
        .unwrap()
}

#[get("/quiz/<quiz_id>")]
fn get_quiz(quiz_id: i32) -> content::RawJson<String> {
    thread::spawn(move || content::RawJson(handlers::get_full_quiz(quiz_id).unwrap()))
        .join()
        .unwrap()
}

#[get("/quiz-stats/<quiz_id>")]
fn get_quiz_stats(quiz_id: i32) -> content::RawJson<String> {
    thread::spawn(move || content::RawJson(handlers::get_stats_quiz(quiz_id).unwrap()))
        .join()
        .unwrap()
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index, get_quiz, get_quiz_stats])
}
