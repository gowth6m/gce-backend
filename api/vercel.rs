use axum::{ routing::{ get, Router }, Json, response::IntoResponse };
use serde::Serialize;
use vercel_runtime::{ process_request, process_response, run_service, Error, ServiceBuilder };

#[derive(Serialize)]
struct Message {
    msg: &'static str,
}

async fn setup_routes() -> Router {
    Router::new().route("/", get(root)).route("/test", get(root2))
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    let app = setup_routes().await;

    let handler = ServiceBuilder::new()
        .map_request(process_request)
        .map_response(process_response)
        .layer(gceBackend::LambdaLayer::default())
        .service(app);

    run_service(handler).await
}

async fn root() -> impl IntoResponse {
    Json(Message { msg: "Hello, World!" })
}

async fn root2() -> impl IntoResponse {
    Json(Message { msg: "Hello, World 2!" })
}
