use axum::{ routing::{ get, post, Router }, Json };
use serde::{ Serialize, Deserialize };
use tracing_subscriber;

#[derive(Serialize)]
struct Message {
    msg: String,
}

#[derive(Deserialize)]
struct InputMessage {
    msg: String,
}

async fn setup_routes() -> Router {
    Router::new().route("/", get(root)).route("/test", get(root2)).route("/test", post(root3))
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let app = setup_routes().await;
    let addr = "0.0.0.0:8080".parse().unwrap();
    axum::Server::bind(&addr).serve(app.into_make_service()).await.unwrap();
}

async fn root() -> Json<&'static str> {
    Json("Hello, World!")
}

async fn root2() -> Json<&'static str> {
    Json("Hello, World 2!")
}

async fn root3(Json(payload): Json<InputMessage>) -> Json<Message> {
    Json(Message { msg: payload.msg })
}
