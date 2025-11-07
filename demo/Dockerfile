FROM eclipse-temurin:21-jdk
WORKDIR /app

COPY gradlew .
COPY gradle ./gradle
COPY build.gradle.kts settings.gradle.kts ./
COPY src ./src

# 한 줄로 통합, clean 대신 rm -rf로 직접 삭제
RUN rm -rf build && chmod +x gradlew && ./gradlew bootJar --no-daemon --no-build-cache -Dorg.gradle.daemon=false -x test && ls -l build/libs

EXPOSE 8080
CMD sh -c 'java -jar build/libs/*.jar'
