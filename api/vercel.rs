use axum::{ routing::{ get, post, Router }, Json };
use vercel_runtime::{ process_request, process_response, run_service, Error, ServiceBuilder };
use serde::{ Serialize, Deserialize };

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
async fn main() -> Result<(), Error> {
    let app = setup_routes().await;

    let handler = ServiceBuilder::new()
        .map_request(process_request)
        .map_response(process_response)
        .layer(gce_backend::LambdaLayer::default())
        .service(app);

    run_service(handler).await
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
