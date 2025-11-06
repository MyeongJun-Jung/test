# 1️⃣ 베이스 이미지: Python + uv 사전 설치
FROM ghcr.io/astral-sh/uv:python3.13-alpine

# 2️⃣ 작업 디렉토리 설정
WORKDIR /app

# 3️⃣ 프로젝트 메타데이터 복사 (dependency install layer)
COPY pyproject.toml uv.lock ./

# 4️⃣ 의존성 설치 (uv 사용)
RUN uv sync --frozen --no-dev

# 5️⃣ 애플리케이션 코드 복사
COPY . .

CMD ["uv", "run", "gunicorn", "main:app", "-w", "3", "-k", "uvicorn.workers.UvicornWorker", "--bind", "0.0.0.0:8000"]
