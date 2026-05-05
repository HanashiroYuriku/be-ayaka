# Ayaka Backend Template

Ayaka is a backend API template developed in the Go programming language, adhering to the principles of Clean Architecture. This system is engineered to provide a scalable, secure, and modular foundation for enterprise-level application development.

## Technical Specifications

* **Language:** [Golang](https://go.dev/)
* **Web Framework:** [Go Fiber v2](https://gofiber.io/)
* **Database:** PostgreSQL
* **ORM:** [GORM](https://gorm.io/)
* **Configuration:** Viper + Godotenv (with Custom Environment Interpolation)
* **Validation:** Go-Playground Validator v10 (with Custom DB Rules)

## Current Features

## Core Functional Features

1. **Strict Layered Architecture:** Implementation of clear boundaries between the `core` layer (Business Logic) and the `adapter` layer (Infrastructure) to maintain framework-agnostic business rules.
2. **Dependency Injection:** Utilization of the Builder Pattern to decouple Handlers, Services, and Repositories, consolidated at the composition root (`builder.go`) for enhanced testability.
3. **Observability and Logging:** Integration of a structured JSON Logger compatible with Datadog/Elasticsearch and automated Request ID tracking for distributed tracing.
4. **Health Monitoring:** A dedicated `/health` endpoint for real-time system status reporting, suitable for Kubernetes and Docker Swarm orchestration.
5. **Schema Management:** Automated database migrations via Go structs (Code-First approach).
6. **Advanced Validation:** Custom database-aware validation tags (`unique`, `incolumn`) designed to enforce referential integrity and mitigate injection risks.
7. **Process Management:** Implementation of graceful shutdown procedures to ensure data integrity during server termination.

## 📂 Project Structure
```text
be-ayaka/
├── cmd/                 # Application entry point (CLI commands, root.go, server.go)
├── config/              # Configuration setup (Viper, Godotenv)
├── internal/            # Private application codebase
│   ├── adapter/         # Infrastructure layer (Database connections, 3rd party APIs)
│   │   ├── database/    # Database initialization (PostgreSQL)
│   │   └── repository/  # GORM Implementations (Fulfills core repository contracts)
│   ├── bootstrap/       # The Wiring (Dependency Injection, App Init, Routes)
│   ├── core/            # Core Business Logic (The Holy Grail - Framework Agnostic)
│   │   ├── entity/      # Business Rules: Pure Data Structs
│   │   ├── port/  # Contracts: Repository Interfaces
│   │   └── service/     # Application Business Rules: Use Cases
│   ├── delivery/        # Transport Mechanism (The Receptionist)
│   │   └── http/        # HTTP Handlers (Fiber controllers: health, user, auth)
│   └── middleware/      # HTTP Interceptors (JWT Auth, Role checking, Request ID)
├── logs/                # Application log files
├── pkg/                 # Reusable, domain-agnostic tools (Hash, JWT, Logger, Validator)
└── .env.sample          # Environment variables template
```

## Development Protocol

Developers are required to follow this sequence when implementing new features:

1. **Domain Entity:** Define your pure data structure in `internal/core/entity/`.
2. **Repository Contract:** Define the database contract in `internal/core/port/`.
3. **Infrastructure Adapter:** Implement the repository interface in `internal/adapter/repository/`.
4. **Service Layer:** Implement the business logic and use cases in `internal/core/service/`.
5. **Delivery Layer:** Implement transport-specific logic in `internal/delivery/http/`.
6. **Composition:** Wire all dependencies (Adapter -> Service -> Handler) together in `internal/bootstrap/builder.go`.
7. **Routing:** Register the endpoint in `internal/bootstrap/routes.go`.

## Usage Examples

### How to Use Custom Validators
This template comes with powerful, database-aware custom validation tags. You can easily apply them to your Data Transfer Object (DTO) structs using the `validate` tag.

**Example Struct:**
```go
package dto

type RegisterUserRequest struct {
    Name     string `json:"name" validate:"required,min=3"`
    
    // unique=table_name->column_name
    // Checks if the email already exists in the 'users' table
    Email    string `json:"email" validate:"required,email,unique=users->email"`
    
    // complexpassword: Enforces uppercase, lowercase, number, and special char (8-12 chars)
    // whitespace: Rejects any space characters
    Password string `json:"password" validate:"required,complexpassword,whitespace"`
    
    // incolumn=table_name->column_name
    // Validates if the provided role_id actually exists in the 'roles' table (Foreign Key check)
    RoleID   string `json:"role_id" validate:"required,incolumn=roles->id"`
}
```

### Available Custom Validation Tags:

* **`unique=[table_name]->[column_name]`**: Returns an error if the value already exists in the specified database table. Perfect for validating unique emails or usernames during registration.
* **`incolumn=[table_name]->[column_name]`**: Returns an error if the value does not exist in the specified database table. Extremely useful for validating Foreign Keys before inserting data.
* **`complexpassword`**: Enforces high-security password policies.
* **`whitespace`**: Ensures the input string contains no spaces.

## Getting Started (Local Development)

### 1. Database Setup
Ensure PostgreSQL is installed and running on your local machine. Create a new, empty database (e.g., `ayaka_db`).

### 2. Environment Configuration (.env)
Clone the .env.sample file to .env in the root directory and configure your credentials:
```env
# server
SERVER_PORT='8000'

#  database configuration
DATABASE_HOST='localhost'
DATABASE_PORT='5432'
DATABASE_USER='postgres'
DATABASE_PASS='password'
DATABASE_NAME='ayaka_db'

# jwt configuration
JWT_SECRET='jwt secret code'
JWT_EXPIRED='24'
```
(Note: Never commit your .env file to a public repository!)

### 3. Install Dependencies
Run the following command to download all required Go modules:
``` bash
go mod tidy
```

### 4. Run the Server
Use the following command to start the server and trigger the auto-migration process:
``` bash
go run main.go svc
```
If the setup is successful, you will see the success log indicating a stable connection to PostgreSQL, followed by the AYAKA ASCII Art in your terminal.

Check the system status by accessing:
```
GET http://localhost:8000/health
```

## Credits & Acknowledgements

* **Hanashiro Yuriku**
* **Gemini** - *AI Pair Programmer*
  Served as a technical consultant and coding assistant, providing architectural insights and collaborative debugging throughout the development journey.

  **Disclaimer:** *The names Ayaka used in this architectural blueprint are inspired by characters from Genshin Impact. This project is open-source, purely for personal/educational purposes, and is not affiliated with, endorsed, or sponsored by HoYoverse/Cognosphere.*