# ---- Build Stage ----
FROM rust:1.82 AS builder

WORKDIR /app
COPY Cargo.toml .
COPY src ./src

# 빌드 캐시 최적화를 위해 deps만 먼저 빌드
RUN cargo fetch
RUN cargo build --release

# ---- Runtime Stage ----
FROM debian:bookworm-slim
WORKDIR /app

# libc 등 최소 런타임 의존성 설치
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# 빌드된 실행 파일 복사
COPY --from=builder /app/target/release/axum_hello .

EXPOSE 8080
CMD ["./axum_hello"]
