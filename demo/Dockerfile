# ---- Build Stage ----
FROM eclipse-temurin:21-jdk AS build
WORKDIR /app

COPY gradlew .
COPY gradle ./gradle
COPY build.gradle.kts settings.gradle.kts ./
COPY src ./src

RUN chmod +x gradlew
RUN ./gradlew clean bootJar --no-daemon -x test && ls -l build/libs

COPY /app/build/libs/*.jar /app/app.jar

EXPOSE 8080
CMD ["java", "-jar", "/app/app.jar"]