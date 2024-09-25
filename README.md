# alert-hub-go

Sample implementation in Golang of an alert management tool using PostgreSQL's pgcrypto.

## Features

- RESTful API for managing alerts
- Secure storage of alert data using PostgreSQL's pgcrypto extension
- Gin web framework for efficient HTTP routing

## Prerequisites

- Docker Compose
- PostgreSQL

## Project Structure

```
alert-hub-go/
├── internal/
│   ├── db/
│   │   └── db.go
│   ├── domain/
│   │   └── alert.go
│   ├── handler/
│   │   └── alert_handler.go
│   ├── repository/
│   │   ├── alert_repository.go
│   │   └── postgres_alert_repository.go
│   └── usecase/
│       └── alert_usecase.go
├── pg-init/
│   └── 01-init.sql
├── Dockerfile
├── compose.yaml
├── go.mod
├── go.sum
└── main.go
```

### Layers of Clean Architecture

1. **Domain Layer** (`internal/domain/`)
   - Contains enterprise-wide business rules and entities
   - Defines the core Alert struct and related interfaces
   - Independent of other layers and external concerns

2. **Use Case Layer** (`internal/usecase/`)
   - Implements application-specific business rules
   - Orchestrates the flow of data to and from entities
   - Contains the AlertUsecase interface and its implementation

3. **Interface Adapters Layer**
   - **Repository** (`internal/repository/`)
     - Implements data access logic
     - Converts data between the format most convenient for entities and use cases, and the format most convenient for external agencies (like databases)
   - **Handler** (`internal/handler/`)
     - Handles HTTP requests and responses
     - Uses the Gin framework to route requests to appropriate use cases

4. **Frameworks and Drivers Layer**
   - **Database** (`internal/db/`)
     - Contains database connection logic
   - **Main** (`main.go`)
     - Ties everything together
     - Sets up dependency injection
     - Initializes the web server

### Benefits of This Architecture

- **Independence of Frameworks**: The core business logic doesn't depend on the existence of a database or a web framework.
- **Testability**: Business rules can be tested without the UI, database, web server, or any other external element.
- **Independence of UI**: The UI can change easily, without changing the rest of the system.
- **Independence of Database**: PostgreSQL could be replaced with another database system without affecting the business rules.
- **Independence of any external agency**: The business rules don't know anything about the outside world.

### Flow of Control

1. HTTP Request → Handler Layer
2. Handler Layer → Use Case Layer
3. Use Case Layer → Repository Layer (if data access is needed)
4. Repository Layer → Database
5. Data flows back through the layers
6. Handler Layer sends HTTP Response

This architecture ensures that inner layers are not dependent on outer layers, maintaining the Dependency Rule of Clean Architecture.

## Setup

Build and run the application using Docker Compose:

```
$ docker-compose up --build
```

## Usage

The API will be available at `http://localhost:8080`. You can use the following endpoints:

- `GET /api/alerts`: Retrieve all alerts
- `POST /api/alerts`: Create a new alert
- `GET /api/alerts/:id`: Retrieve a specific alert
- `PATCH /api/alerts/:id`: Update an existing alert

Example of creating a new alert:

```bash
$ curl -X POST http://localhost:8080/api/alerts \
  -H "Content-Type: application/json" \
  -d '{
    "subject": "Test Alert",
    "body": "This is a test alert.",
    "identifier": "TEST-001",
    "urgency": "MEDIUM",
    "status": "OPEN"
  }'
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
