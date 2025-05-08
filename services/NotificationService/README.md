
# 📘 Notification Microservice

## 📖 Overview
This microservice handles **user notifications** in a Library Management System. It listens to **Apache Kafka** topics for events like:

- Due date reminders
- Overdue alerts
- Availability of reserved books

These events are persisted to a **relational database**, and **REST APIs** are exposed to fetch, create, delete, or update notification statuses.

---

## 🏗️ Architecture & Technology Stack

| Component        | Technology             |
|------------------|-------------------------|
| Framework        | Spring Boot             |
| Messaging        | Apache Kafka            |
| Database         | JPA (MySQL)             |
| Validation       | Jakarta Bean Validation |
| Serialization    | Jackson (JSON)          |
| API Layer        | Spring MVC (RESTful)    |

---

## 🧱 Package Structure

```
com.library.management
├── config              # Kafka Consumer Configuration
├── controller          # REST Controller
├── entity              # JPA Entity and Error Representation
├── exceptions          # Global Exception Handling
├── service             # Business Logic Layer
├── repo                # Data Access Layer (assumed present)
```

---

## 🛠️ Setup Instructions

### ✅ Prerequisites

- Java 17+
- Maven
- MySQL Database
- Apache Kafka with Zookeeper

---

### 🔧 Steps to Run

1. **Clone the repository**
   ```bash
   git clone <repo-url>
   cd library-management/services/NotificationService
   ```

2. **Configure your database and Kafka in**
   `src/main/resources/application.properties`:
   ```properties
   spring.datasource.url=jdbc:mysql://localhost:3306/librarydb
   spring.datasource.username=root
   spring.datasource.password=your_password
   spring.jpa.hibernate.ddl-auto=update

   spring.kafka.bootstrap-servers=localhost:9092
   spring.kafka.consumer.group-id=notification-group
   ```

3. **Start Kafka and Zookeeper**
   ```bash
   # Start Zookeeper
   bin/zookeeper-server-start.sh config/zookeeper.properties

   # Start Kafka
   bin/kafka-server-start.sh config/server.properties
   ```

4. **Create Kafka Topics**
   ```bash
   bin/kafka-topics.sh --create --topic book.due_date_reminder --bootstrap-server localhost:9092
   bin/kafka-topics.sh --create --topic book.overdue --bootstrap-server localhost:9092
   bin/kafka-topics.sh --create --topic book.available_reserved --bootstrap-server localhost:9092
   ```

5. **Build the project**
   ```bash
   ./mvnw clean install
   ```

6. **Run the application**
   ```bash
   ./mvnw spring-boot:run
   ```

---

## 🔁 Kafka Consumer

- **Topics Consumed**:
  - `book.due_date_reminder`
  - `book.overdue`
  - `book.available_reserved`

- **Group ID**: `notification-group`

- **Behavior**:
  - Consumes JSON messages
  - Extracts `userId` and `message`
  - Persists into the `notifications` table

---

## 📡 REST APIs

### `GET /notifications/{userId}`
Get a paginated list of notifications for a user.

- **Query Params**:
  - `page` (default: 0)
  - `size` (default: 10)

- **Response**: `Page<Notification>`

---

### `POST /notifications`

Create a new notification.

#### Request Body:
```json
{
  "userId": 123,
  "message": "Your reserved book is now available."
}
```

- **Response**: `Notification`

---

### `DELETE /notifications/{id}`

Delete a notification by ID.

- **Response**: HTTP 200 OK

---

### `PATCH /notifications/{id}/seen`

Mark a notification as seen.

- **Response**: HTTP 200 OK

---

## 📂 Entity: Notification

| Field     | Type           | Description                     |
|-----------|----------------|---------------------------------|
| id        | Long           | Primary key                     |
| userId    | Long           | User receiving the notification |
| message   | String         | Notification content            |
| timestamp | LocalDateTime  | Creation time                   |
| seen      | boolean        | Read status                     |

---

## ❗ Error Handling

- **Validation Errors**: Returns 400 with a list of error messages.
- **Generic Errors**: Returns 500 with error details.

---

## 📌 Best Practices Followed

- Clean separation of concerns (controller, service, consumer)
- Centralized error handling via `@RestControllerAdvice`
- Validation with Jakarta annotations
- Kafka-based async event-driven communication
- Pagination support for efficient data access

---

## 🚀 Future Enhancements

- Add authentication & authorization
- Support for email/SMS/push notifications
- Caching for frequently accessed notifications
- Configurable message templates

---

## 🧪 Sample Test Data (SQL)

Run the following queries in MySQL Workbench or CLI:

```sql
-- User ID 1
INSERT INTO notifications (user_id, message, timestamp, seen) VALUES
(1, 'Your book "Clean Code" is due in 2 days.', NOW(), false),
(1, 'Your reserved book "Domain-Driven Design" is now available.', NOW() - INTERVAL 1 DAY, false),
(1, 'Reminder: Return "Design Patterns" today to avoid overdue.', NOW() - INTERVAL 2 DAY, true);

-- User ID 2
INSERT INTO notifications (user_id, message, timestamp, seen) VALUES
(2, 'Overdue alert: "Refactoring" was due 3 days ago.', NOW() - INTERVAL 3 DAY, false),
(2, 'New book "Microservices Patterns" added to library.', NOW(), true);

-- User ID 3
INSERT INTO notifications (user_id, message, timestamp, seen) VALUES
(3, 'You have successfully renewed "Effective Java".', NOW() - INTERVAL 5 HOUR, true),
(3, 'Upcoming due date: "Kubernetes in Action".', NOW() - INTERVAL 1 HOUR, false);
```
