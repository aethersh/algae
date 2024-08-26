use api::{make_bgp_router, make_traceroute_router};
use axum::{
    http::StatusCode,
    response::IntoResponse,
    routing::{get, post},
    Json, Router,
};
use serde::Deserialize;
mod api;
mod mtr;
use std::net::IpAddr;

#[derive(Debug, Deserialize)]
struct PingRequest {
    ip: IpAddr,
}

async fn ping_handler(Json(input): Json<PingRequest>) -> impl IntoResponse {
    let ip_str = input.ip.to_string();
    if input.ip.is_ipv4() {
        println!("Pinging IPv4 address: {}", ip_str);
        (StatusCode::OK, Json("{ message: 'IPv4 Address Accepted' }"))
    } else if input.ip.is_ipv6() {
        println!("Pinging IPv6 address: {}", ip_str);
        (StatusCode::OK, Json("{ message: 'IPv6 Address Accepted' }"))
    } else {
        println!("Invalid IP address");
        (
            StatusCode::BAD_REQUEST,
            Json("{ message: 'Invalid IP address' }"),
        )
    }
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/", get(|| async { "Alive!" }))
        // .merge(make_ping_router())
        .route("/ping", post(ping_handler))
        .merge(make_bgp_router())
        .merge(make_traceroute_router())    // .route("/test", post(post_handle_test))
    ;

    // run our app with hyper, listening globally on port 2152
    let listener = tokio::net::TcpListener::bind("0.0.0.0:2152").await.unwrap();
    println!("Server running on port 2152");
    axum::serve(listener, app).await.unwrap();
}

// async fn post_handle_test(r: Request<Body>) -> impl IntoResponse {
//     let body = r.body();

//     StatusCode::BAD_REQUEST
// }
