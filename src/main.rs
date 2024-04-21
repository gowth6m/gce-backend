use axum::{ routing::{ get }, Router };
use tokio::net::TcpListener;

mod handlers;
use handlers::intro::intro;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let app = Router::new().route("/", get(intro));

    let listener = TcpListener::bind("0.0.0.0:8080").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
