use axum::{routing::get, Json, Router};
use serde_json::json;

async fn hello() -> Json<serde_json::Value> {
    Json(json!({ "hello": "world" }))
}

#[tokio::main]
async fn main() {
    // 라우터 구성
    let app = Router::new().route("/", get(hello));

    // 0.7 스타일: TcpListener + axum::serve
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080")
        .await
        .expect("failed to bind");

    println!("🚀 Server running on http://0.0.0.0:8080");

    axum::serve(listener, app)
        .await
        .expect("server error");
}
