use axum::extract::Path;

use ping::ping;
use rand::random;
use std::{net::IpAddr, time::Duration};
const _PING_TIMEOUT: Duration = Duration::from_secs(1);

// Axum handler for POST /ping/:ip
pub async fn _ping_ip(Path(ip): Path<String>) -> String {
    let ip = ip.parse::<IpAddr>().unwrap();
    let _ping_session = ping(
        ip,
        Some(_PING_TIMEOUT),
        Some(166),
        Some(3),
        Some(5),
        Some(&random()),
    )
    .unwrap();

    "ping".to_string()
}
