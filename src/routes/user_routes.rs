use axum::{ Router, routing::{ get, post }, Json };
use serde::{ Deserialize, Serialize };

use crate::services::user::user_handlers::get_current_user;

pub fn user_routes() -> Router {
    Router::new()
        .route("/", get(user_fn))
        .route("/current", get(get_current_user))
        .route("/", post(user_post_fn))
}

async fn user_fn() -> Json<&'static str> {
    Json("Hello, User!")
}

async fn user_post_fn(Json(payload): Json<InputMessage>) -> Json<Message> {
    Json(Message { msg: payload.msg })
}

#[derive(Deserialize)]
struct InputMessage {
    msg: String,
}

#[derive(Serialize)]
struct Message {
    msg: String,
}
