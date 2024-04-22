use gce_backend::setup_routes;
use tracing_subscriber;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let app = setup_routes().await;
    let addr = "0.0.0.0:8080".parse().unwrap();
    axum::Server::bind(&addr).serve(app.into_make_service()).await.unwrap();
}
