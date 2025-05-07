# Build stage
FROM maven:3.8-openjdk-17 AS build
WORKDIR /app

# Copy the Maven wrapper files
COPY mvnw .
COPY .mvn .mvn

# Make the mvnw script executable
RUN chmod +x mvnw

# Copy the project files
COPY pom.xml .
COPY src src

# Build the application
RUN ./mvnw clean install -DskipTests

# Runtime stage
FROM amazoncorretto:17-alpine
WORKDIR /app

# Create a non-root user to run the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Copy the jar from the build stage
COPY --from=build /app/target/*.jar app.jar

# Set environment variables
ENV SPRING_PROFILES_ACTIVE=prod

# Expose the application port
EXPOSE 8085

# Run the application
ENTRYPOINT ["java", "-jar", "app.jar"]