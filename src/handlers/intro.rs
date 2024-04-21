use axum::Json;
use serde::Serialize;

#[derive(Serialize)]
pub struct ApiResponse {
    pub code: u16,
    pub message: String,
}

pub async fn intro() -> Json<ApiResponse> {
    Json(ApiResponse {
        code: 200,
        message: "hello world!".to_string(),
    })
}