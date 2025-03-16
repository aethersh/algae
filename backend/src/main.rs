use api::{make_bgp_router, make_traceroute_router};
use axum::{
    routing::{get, post},
    Router,
};
mod api;
mod mtr;
mod ping;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(|| async { "Alive!" }))
        // .merge(make_ping_router())
        .route("/ping", post(ping::ping_handler))
        .merge(make_bgp_router())
        .merge(make_traceroute_router());

    // run our app with hyper, listening globally on port 2152
    let listener = tokio::net::TcpListener::bind("0.0.0.0:2152").await.unwrap();
    println!("Server running on port 2152");
    axum::serve(listener, app).await.unwrap();
}
