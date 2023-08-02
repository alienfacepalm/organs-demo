// main.rs

use firestore::firestore_client::FirestoreClient;
use firestore::Organ;
use serde::Deserialize;
use std::env;
use std::net::SocketAddr;
use warp::http::Response;
use warp::{Filter, Rejection, Reply};

#[derive(Debug, Deserialize)]
struct FirestoreConfig {
    project_id: String,
    api_key_file: String,
}

impl FirestoreConfig {
    fn from_env() -> Result<Self, Box<dyn std::error::Error>> {
        let api_key_file = env::var("../../organs-demo-api-key.json")?;
        let project_id = env::var("organs-demo")?;

        Ok(FirestoreConfig {
            project_id,
            api_key_file,
        })
    }
}

#[tokio::main]
async fn main() {
    // Load Firestore config
    let firestore_config =
        firestore_config::FirestoreConfig::from_env().expect("Failed to load Firestore config");

    // Initialize Firestore client
    let firestore_client =
        FirestoreClient::new(&firestore_config.project_id, &firestore_config.api_key_file)
            .expect("Failed to create Firestore client");

    // Create the filter chain
    let get_organs = warp::get()
        .and(warp::path("organs"))
        .and(with_firestore_client(firestore_client.clone()))
        .and_then(handle_get_organs);

    let routes = get_organs;

    // Start the server
    let addr: SocketAddr = "127.0.0.1:8080".parse().expect("Invalid address");
    warp::serve(routes).run(addr).await;
}

async fn handle_get_organs(
    firestore_client: FirestoreClient,
) -> Result<impl Reply, Rejection> {
    let organs = firestore_client
        .get_organs()
        .await
        .map_err(|e| warp::reject::custom(FirestoreError(e)))?;
    Ok(Response::json(&organs))
}

fn with_firestore_client(
    client: FirestoreClient,
) -> impl Filter<Extract = (FirestoreClient,), Error = std::convert::Infallible> + Clone {
    warp::any().map(move || client.clone())
}
