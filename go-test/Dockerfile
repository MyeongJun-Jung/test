# 1️⃣ Build stage
FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod ./
COPY . .
RUN go mod tidy

# ✅ 명시적으로 빌드 결과 경로 지정
RUN go build -o /app/server main.go

# 2️⃣ Runtime stage
FROM debian:bookworm-slim
WORKDIR /app

# ✅ 정확한 빌드 결과 복사
COPY --from=builder /app/server /app/server

# ✅ 실행 권한 보장 (보통 Go 빌드 시 자동이지만 확실히)
RUN chmod +x /app/server

EXPOSE 8080
CMD ["/app/server"]
