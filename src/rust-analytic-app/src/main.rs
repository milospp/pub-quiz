// #[macro_use]
// extern crate rocket;

// mod utils;

// use rocket::serde::json::Json;
// use rocket::{routes, Build, Rocket};
// use std::thread;
// use utils::Quiz;

// #[get("/<id>")]
// fn get_all_active_quizes(id: i32) -> Json<String> {
//     thread::spawn(move || Json(utils::get_qll_quizzes(id).unwrap()))
//         .join()
//         .unwrap()
// }

// #[launch]
// fn rocket() -> Rocket<Build> {
//     rocket::build().mount("/api/analytics", routes![get_all_active_quizes])
// }

#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index])
}