use axum::{ response::IntoResponse, Json };

use serde::{ Serialize, Deserialize };

pub async fn get_current_user() -> impl IntoResponse {
    let user = User {
        name: "John Doe".to_string(),
        address: "123 Main St\nSpringfield, IL 62701".to_string(),
    };

    Json(user)
}

#[derive(Serialize, Deserialize)]
struct User {
    name: String,
    address: String,
}
