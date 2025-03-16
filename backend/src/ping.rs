use axum::{response::IntoResponse, Json};
use fastping_rs::PingResult::{Idle, Receive};
use fastping_rs::Pinger;
use serde::{Deserialize, Serialize};
use std::net::IpAddr;
use std::time::Duration;

// HTTP Request
#[derive(Debug, Deserialize)]
pub struct PingRequest {
    ip: IpAddr,
}

// HTTP Response
#[derive(Debug, Serialize)]
pub struct PingResponse {
    ip: IpAddr,
    pings: Vec<SerializablePingResult>,
    max_rtt: f64,
    min_rtt: f64,
    avg_rtt: f64,
    std_dev_rtt: f64,
    total_count: u64,
    lost_count: u64,
}

impl PingResponse {
    pub fn make_ping_response(ip_addr: IpAddr, pings: Vec<SerializablePingResult>) -> Self {
        let mut total_rtt_micros = 0 as u128;
        let mut min_rtt_micros = std::u128::MAX;
        let mut max_rtt_micros = 0 as u128;
        let mut lost_count = 0;

        for ping in &pings {
            match ping.rtt {
                Some(rtt) => {
                    let rtt_micros = rtt.as_micros();
                    total_rtt_micros += rtt_micros;
                    min_rtt_micros = min_rtt_micros.min(rtt_micros);
                    max_rtt_micros = max_rtt_micros.max(rtt_micros);
                }
                None => lost_count += 1,
            }
        }

        let total_count = pings.len() as u64;
        // Calculate the average RTT microseconds
        let avg_rtt_micros = total_rtt_micros as f64 / total_count as f64;

        // Calculate the standard deviation of the RTT microseconds using the average
        let std_dev = pings
            .iter()
            .filter_map(|ping| ping.rtt)
            .map(|rtt| (rtt.as_micros() as f64 - avg_rtt_micros).powi(2))
            .sum::<f64>()
            .sqrt()
            / total_count as f64;

        Self {
            ip: ip_addr,
            pings,
            max_rtt: max_rtt_micros as f64 / 1000.0 as f64,
            min_rtt: min_rtt_micros as f64 / 1000.0 as f64,
            avg_rtt: avg_rtt_micros as f64 / 1000.0 as f64,
            std_dev_rtt: std_dev,
            total_count,
            lost_count,
        }
    }
}

// Intermediate struct to serialize the ping results
#[derive(Debug, Serialize)]
pub struct SerializablePingResult {
    pub success: bool,
    pub ip: IpAddr,
    pub rtt: Option<Duration>,
    pub seq: u64,
}

impl SerializablePingResult {
    fn from_ping_result(result: fastping_rs::PingResult, seq: u64) -> Self {
        match result {
            Idle { addr } => Self {
                success: false,
                ip: addr,
                rtt: None,
                seq,
            },
            Receive { addr, rtt } => Self {
                success: true,
                ip: addr,
                rtt: Some(rtt),
                seq,
            },
        }
    }
}

const MAX_PINGS: u64 = 10;

pub async fn ping_handler(Json(input): Json<PingRequest>) -> impl IntoResponse {
    let ip_str = input.ip.to_string();

    let (pinger, results) = match Pinger::new(Some(2), Some(56)) {
        Ok((pinger, results)) => (pinger, results),
        Err(e) => panic!("Error creating pinger: {}", e),
    };

    // Start pinger with the given IP address
    pinger.add_ipaddr(&ip_str);
    pinger.run_pinger();

    let mut ping_results: Vec<SerializablePingResult> = Vec::new();

    // Collect 10 pings
    for seq in 0..MAX_PINGS {
        match results.recv() {
            Ok(result) => ping_results.push(SerializablePingResult::from_ping_result(result, seq)),
            Err(_) => panic!("Worker threads disconnected before the solution was found!"),
        }
    }

    // Stop the pinger
    pinger.stop_pinger();

    // Return ping response
    PingResponse::make_ping_response(input.ip, ping_results);
}

// pub async fn ping_stream_handler(
//     Json(input): Json<PingRequest>,
// ) -> Sse<impl Stream<Item = Result<Event, Infallible>>> {
//     let ip_str = input.ip.to_string();

//     let (pinger, results) = match Pinger::new(None, Some(56)) {
//         Ok((pinger, results)) => (pinger, results),
//         Err(e) => panic!("Error creating pinger: {}", e),
//     };

//     let ping_stream =

//     Sse::new(ping_stream).keep_alive(
//         axum::response::sse::KeepAlive::new()
//             .interval(Duration::from_secs(1))
//             .text("keep-alive-text"),
//     )
// }
