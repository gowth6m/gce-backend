use axum::{ Router, routing::get, Json };

pub fn root_routes() -> Router {
    Router::new().route("/", get(root_fn))
}

async fn root_fn() -> Json<&'static str> {
    Json("Hello, World!")
}
