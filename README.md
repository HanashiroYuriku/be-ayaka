# 🌸 Ayaka Backend Template

A Golang (Go) Backend API template designed with Clean Architecture. This project focuses on scalability, security, and seamless team collaboration.

## 🚀 Core Technologies

* **Language:** [Golang](https://go.dev/)
* **Web Framework:** [Go Fiber v2](https://gofiber.io/)
* **Database:** PostgreSQL
* **ORM:** [GORM](https://gorm.io/)
* **Configuration:** Viper + Godotenv (with Custom Environment Interpolation)
* **Validation:** Go-Playground Validator v10 (with Custom DB Rules)

## ✨ Current Features

1.  **Modular Architecture:** Clear separation of layers between `adapter`, `domain`, and `core` to maintain clean boundaries.
2.  **Advanced Configuration System:** Supports environment variable interpolation (e.g., `${DATABASE_HOST:localhost}`) paired with a Fail-Fast mechanism for secure deployments.
3.  **Database Auto-Migration (Code-First):** Database tables and relationships are automatically generated via Golang structs using GORM.
4.  **Enterprise Custom Validator:**
    * Human-readable error messages (automatically translated to English).
    * `unique`: Checks for data duplication directly in the database (Anti SQL-Injection).
    * `incolumn`: Ensures valid data/ID relationships in the database.
    * `complexpassword`: Strict password security validation.
    * `whitespace`: Rejects whitespaces in strict input fields.
5.  **Graceful Shutdown:** Safely terminates the server and cleans up resources without dropping ongoing requests.
6.  **Custom JSON Logger:** Structured logging using JSON format for easy integration with observability and monitoring systems (like Elasticsearch or Datadog).

## 📂 Project Structure
```
├── cmd/            # Application entry point
├── config/         # Viper & Godotenv configurations
├── internal/       # Core business logic (Adapter, Domain, Core)
└── pkg/            # Reusable packages (Logger, Validator)
```

## 💡 Usage Examples

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

* **`unique=<table_name>-><column_name>`**: Returns an error if the value already exists in the specified database table. Perfect for validating unique emails or usernames during registration.
* **`incolumn=<table_name>-><column_name>`**: Returns an error if the value does not exist in the specified database table. Extremely useful for validating Foreign Keys before inserting data.
* **`complexpassword`**: Enforces high-security password policies.
* **`whitespace`**: Ensures the input string contains no spaces.

## 🛠️ Getting Started (Local Development)

### 1. Database Setup
Ensure PostgreSQL is installed and running on your local machine. Create a new, empty database (e.g., `ayaka_db`).

### 2. Environment Configuration (.env)
Create a `.env` file in the root directory (alongside `main.go`) and configure your database credentials:

```env
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASS='your_super_secret_password'
DATABASE_NAME=ayaka_db
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
If the setup is successful, you will see the success log indicating a stable connection to PostgreSQL, followed by the AYAKA ASCII Art in your terminal. The server will be accessible at http://localhost:8000.

## 🌟 Credits & Acknowledgements

* **Hanashiro Yuriku**
* **Gemini** - *AI Pair Programmer*
  Served as a technical consultant and coding assistant, providing architectural insights and collaborative debugging throughout the development journey.