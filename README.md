# Task Management Service (Go Microservice)

This is a simple **Task Management Microservice** built in **Go** as part of the Alle Backend Assignment.

The service allows users to:

✅ Create Tasks  
✅ Read Tasks (single + list with pagination + filtering)  
✅ Update Tasks  
✅ Delete Tasks  

It demonstrates:

✅ Microservices architecture principles  
✅ Single Responsibility Principle  
✅ Clean RESTful API Design  
✅ Scalability  
✅ Future readiness for inter-service communication  
✅ API documentation via Swagger UI  

---

## Features

- CRUD APIs for Task entity
- Pagination on `GET /v1/tasks` → supports `page` and `size`
- Filtering on `GET /v1/tasks` → supports `status=Pending|InProgress|Completed`
- Well-structured Go project → SRP at each layer
- RESTful API Design → versioned `/v1/tasks`
- Swagger UI → self-documenting API → http://localhost:8080/swagger/index.html
- Stateless → horizontally scalable

---

## Project Structure

```txt
task-service/
├── cmd/task-service/main.go         # Entry point
├── internal/api/                    # HTTP Handlers (Controllers)
├── internal/service/                # Business logic
├── internal/repository/             # DB layer
├── internal/model/                  # Domain model (Task)
├── internal/db/                     # DB connection setup
├── docs/                            # Swagger generated docs
├── Dockerfile                       # (optional) for containerization
├── docker-compose.yml               # (optional) for DB container
├── go.mod / go.sum                  # Dependencies
└── README.md
```

---

## How to Run & Test the Service

### Prerequisites

- Go 1.22+ installed → https://go.dev/dl/
- MySQL running locally on port 3306 (or in Docker)
- Swagger CLI installed (for regenerating docs if needed):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

### 1️⃣ Setup Database

In MySQL:

```sql
CREATE DATABASE taskdb;
```

Or use an existing database and update the config in `internal/db/db.go`:

```go
host := "localhost"
port := "3306"
user := "user"          // your MySQL username
password := "password"  // your MySQL password
dbname := "taskdb"
```

---

### 2️⃣ Install Go dependencies

```bash
go mod tidy
```

---

### 3️⃣ Generate Swagger docs (only needed if you change Swagger comments)

```bash
~/go/bin/swag init -g cmd/task-service/main.go
```

---

### 4️⃣ Run the application

```bash
go run cmd/task-service/main.go
```

---

### 5️⃣ Access the API

The server will start at:

```txt
http://localhost:8080
```

Swagger UI available at:

```txt
http://localhost:8080/swagger/index.html
```

---

### 6️⃣ Example API testing

#### List tasks

```http
GET http://localhost:8080/v1/tasks?page=1&size=10&status=Pending
```

#### Create a task

```http
POST http://localhost:8080/v1/tasks
Content-Type: application/json

{
  "title": "Test my first task",
  "description": "Trying the POST API",
  "status": "Pending",
  "due_date": "2025-12-31T00:00:00Z"
}
```

#### Update a task

```http
PUT http://localhost:8080/v1/tasks/1
Content-Type: application/json

{
  "title": "Updated task",
  "description": "Updated description",
  "status": "Completed",
  "due_date": "2025-12-31T00:00:00Z"
}
```

#### Delete a task

```http
DELETE http://localhost:8080/v1/tasks/1
```

---

## Demonstration of Microservices Concepts

### Single Responsibility Principle

- **model** → defines domain model → Task
- **repository** → encapsulates DB access
- **service** → contains business logic
- **api** → exposes HTTP endpoints
- **cmd** → main entry point to wire dependencies

Each layer has a **single clear responsibility**.

---

### API Design

- RESTful principles:
  - `/v1/tasks` → resource-based
  - Versioning → `/v1` → future-proof
  - Consistent naming → standard HTTP verbs
- Pagination → `page`, `size` query params
- Filtering → `status` query param
- Swagger UI → fully documented

---

### Scalability

- Service is **stateless** → horizontal scaling supported:
  - No local session state
  - Multiple instances can be deployed behind a load balancer
- DB connection pooling via GORM
- Cloud-ready architecture → easily containerized via Docker

Example horizontal scaling:

```txt
             +-------------------+
             | Load Balancer (LB) |
             +---------+---------+
                       |
        +--------------+--------------+
        |                             |
 +------------+                +------------+
 | App Instance|                | App Instance|
 +------------+                +------------+
                       |
                   MySQL DB
```

---

### Inter-Service Communication

If a **User Service** were added, we could integrate via:

1️⃣ **REST API**  
→ Simple synchronous calls (GET /users/{id})

2️⃣ **gRPC**  
→ High-performance, typed communication

3️⃣ **Message Queue (Kafka, RabbitMQ)**  
→ Event-driven architecture:
- Emit `TaskCreated` or `TaskDeleted` events
- Other services (User Service) can consume these asynchronously
- Enables loose coupling & high scalability

---

## Conclusion

✅ The service demonstrates **solid microservices architecture**:

- Clean separation of concerns → SRP  
- RESTful API Design → versioned, documented, consistent  
- Scalability → stateless, horizontally scalable  
- Inter-service communication ready → REST/gRPC/MQ patterns supported  

---

## Author

Satyam Lal  
GitHub: [pvnptl](https://github.com/pvnptl)  
Assignment for: **Alle | Backend Assignment - Tasks**

---

# ✅ Done!
