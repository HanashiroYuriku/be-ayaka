# 🌸 Ayaka Backend Template

A Golang (Go) Backend API template designed with Clean Architecture. This project focuses on scalability, security, and seamless team collaboration.

## Core Technologies

* **Language:** [Golang](https://go.dev/)
* **Web Framework:** [Go Fiber v2](https://gofiber.io/)
* **Database:** PostgreSQL
* **ORM:** [GORM](https://gorm.io/)
* **Configuration:** Viper + Godotenv (with Custom Environment Interpolation)
* **Validation:** Go-Playground Validator v10 (with Custom DB Rules)

## Current Features

1. **Strict Layered Architecture:** Clear separation between `core` (Entities, Interfaces, Business Logic) and `adapter` (Database implementation, External APIs) to ensure the business rules remain agnostic of frameworks.
2. **Dependency Injection (Builder Pattern):** Handlers, Services, and Repositories are decoupled and wired cleanly at the composition root (`builder.go`), making the app 100% testable.
3. **Advanced Observability & Logging:** * Custom JSON Logger ready for Datadog/Elasticsearch.
   * **Request ID:** Automatically tracks request flows across logs and responses for painless debugging.
4. **Enterprise Health Check:** A `/health` endpoint that actively pings the database connection to report real-time system status (Up/Down) for Kubernetes/Docker Swarm load balancers.
5. **Database Auto-Migration (Code-First):** Tables and relationships are automatically generated via Golang structs.
6. **Smart Custom Validators:** Database-aware validation rules (`unique`, `incolumn`) to prevent SQL Injection and enforce foreign key constraints seamlessly.
7. **Graceful Shutdown:** Safely terminates the server and cleans up database connections without dropping ongoing client requests.

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
│   │   ├── repository/  # Contracts: Repository Interfaces
│   │   └── service/     # Application Business Rules: Use Cases
│   ├── delivery/        # Transport Mechanism (The Receptionist)
│   │   └── http/        # HTTP Handlers (Fiber controllers: health, user, auth)
│   └── middleware/      # HTTP Interceptors (JWT Auth, Role checking, Request ID)
├── logs/                # Application log files
├── pkg/                 # Reusable, domain-agnostic tools (Hash, JWT, Logger, Validator)
└── .env.sample          # Environment variables template
```

## Development Workflow

To maintain the integrity of Clean Architecture, follow this sequence when adding a new feature:

1. **Entity (The Corest):** Define your pure data structure in `internal/core/entity/`.
2. **Repository Interface (The Contract):** Define the database contract in `internal/core/repository/`.
3. **Repository Adapter (The Implementation):** Write the GORM SQL logic in `internal/adapter/repository/` to fulfill the contract.
4. **Service (The Brain):** Implement the business logic and use cases in `internal/core/service/`.
5. **Delivery/Handler (The Receptionist):** Create the HTTP delivery logic in `internal/delivery/http/`.
6. **Builder (The Matchmaker):** Wire all dependencies (Adapter -> Service -> Handler) together in `internal/bootstrap/builder.go`.
7. **Routes (The Door):** Register your new endpoint to the Fiber app in `internal/bootstrap/routes.go`.

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