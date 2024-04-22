use gce_backend::setup_routes;
use vercel_runtime::{ process_request, process_response, run_service, Error, ServiceBuilder };

mod lambda;

#[tokio::main]
async fn main() -> Result<(), Error> {
    let app = setup_routes().await;

    let handler = ServiceBuilder::new()
        .map_request(process_request)
        .map_response(process_response)
        .layer(lambda::LambdaLayer::default())
        .service(app);

    run_service(handler).await
}
