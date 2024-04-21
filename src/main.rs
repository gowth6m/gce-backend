use axum::{ routing::{ get }, Router, response::{ Json } };
use tokio::net::TcpListener;

async fn setup_routes() -> Router {
    Router::new().route("/", get(root)).route("/test", get(root2))
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let app = setup_routes().await;

    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn root() -> Json<&'static str> {
    Json("Hello, World!")
}

async fn root2() -> Json<&'static str> {
    Json("Hello, World 2!")
}
