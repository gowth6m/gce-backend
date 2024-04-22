use axum::Router;

pub mod routes;
pub mod services;

pub async fn setup_routes() -> Router {
    Router::new()
        .nest("/", Router::new().merge(routes::root_routes::root_routes()))
        .nest("/user", Router::new().merge(routes::user_routes::user_routes()))
}
