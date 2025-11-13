use axum::{
    extract::ws::{Message, WebSocket, WebSocketUpgrade},
    extract::State,
    response::IntoResponse,
    routing::get,
    Json, Router,
};
use futures::{SinkExt, StreamExt};
use serde_json::{json, Value};
use tokio::time::{sleep, Duration};
use std::sync::Arc;

async fn hello() -> Json<Value> {
    Json(json!({ "hello": "world" }))
}

// ---- WebSocket Handler ----
async fn ws_handler(ws: WebSocketUpgrade, State(_state): State<AppState>) -> impl IntoResponse {
    println!("⚡ WebSocket connection requested");
    ws.on_upgrade(handle_socket)
}

async fn handle_socket(mut socket: WebSocket) {
    println!("✅ Client connected");

    while let Some(Ok(msg)) = socket.next().await {
        if let Message::Text(text) = msg {
            let parsed: Value = match serde_json::from_str(&text) {
                Ok(v) => v,
                Err(_) => {
                    let _ = socket
                        .send(Message::Text(json!({"error": "invalid JSON"}).to_string()))
                        .await;
                    continue;
                }
            };

            // Expecting ["join_ref","ref","topic","event","payload"]
            if !parsed.is_array() || parsed.as_array().unwrap().len() < 4 {
                let _ = socket
                    .send(Message::Text(json!({"error": "invalid format"}).to_string()))
                    .await;
                continue;
            }

            let arr = parsed.as_array().unwrap();
            let topic = arr[2].as_str().unwrap_or("");
            let event = arr[3].as_str().unwrap_or("");

            match event {
                "phx_join" => {
                    let resp = json!(["1", "1", topic, "phx_reply", { "status": "ok" }]);
                    let _ = socket.send(Message::Text(resp.to_string())).await;
                    println!("📡 JOIN {}", topic);
                }
                "ping" => {
                    let resp =
                        json!(["2", "2", topic, "pong", { "msg": "hello" }]);
                    let _ = socket.send(Message::Text(resp.to_string())).await;

                    // Optional small delay to avoid too-fast echo
                    sleep(Duration::from_millis(1)).await;
                }
                _ => {
                    let resp = json!(["?", "?", topic, "unknown_event"]);
                    let _ = socket.send(Message::Text(resp.to_string())).await;
                }
            }
        }
    }

    println!("❌ Client disconnected");
}

// Global state struct (expandable later)
#[derive(Clone)]
struct AppState {}

#[tokio::main]
async fn main() {
    let state = AppState {};

    let app = Router::new()
        .route("/", get(hello))
        .route("/ws", get(ws_handler))
        .with_state(state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080")
        .await
        .expect("failed to bind");

    println!("🚀 Rust WebSocket Benchmark Server running at ws://0.0.0.0:8080/ws");

    axum::serve(listener, app)
        .await
        .expect("server error");
}
