use axum::{ routing::get, Router };
use vercel_runtime::{ process_request, process_response, run_service, Error, ServiceBuilder };

mod handlers;
use handlers::intro::intro;

#[tokio::main]
async fn main() -> Result<(), Error> {
    // run(handler).await

    let app = Router::new().route("/", get(intro));

    let handler = ServiceBuilder::new()
        .map_request(process_request)
        .map_response(process_response)
        .service(app);

    run_service(handler).await
}
