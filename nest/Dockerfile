
# Use the official Node.js image as the base image
FROM node:24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install the application dependencies
RUN npm ci --omit=dev

# Copy the rest of the application files
COPY . .

# Build the NestJS application
RUN npm run build

# 5️⃣ pm2 글로벌 설치
RUN npm install -g pm2

# 6️⃣ 서비스 포트 노출
EXPOSE 3000

# 7️⃣ pm2-runtime으로 NestJS 실행 (클러스터 모드 자동 CPU 수 감지)
CMD ["pm2-runtime", "dist/main.js", "-i", "max", "--name", "nestjs-app"]