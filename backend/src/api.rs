use axum::{routing::get, routing::post, Router};
// use birdc::Client;

/*  */

// --------- Ping Stuff ---------
// pub fn make_ping_router() -> Router {
//     let ping_router = Router::new()
//         .route("/ping", post(|| async { "ponged" }))
//         .route("/ping/:ip", get(|| async { "pong" }));

//     ping_router
// }

// --------- Traceroute Stuff ---------
pub fn make_traceroute_router() -> Router {
    let trace_router = Router::new().route("/traceroute", post(|| async { "traced" }));

    trace_router
}

// --------- BGP Stuff ---------
pub fn make_bgp_router() -> Router {
    let bgp_router = Router::new()
        .route(
            "/bgp",
            get(|| async { "bordered gatewayed and protocoled" }),
        )
        .route("/bgp/:ip", get(|| async { "pong" }));

    bgp_router
}
