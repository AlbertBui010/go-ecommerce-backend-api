# Go E-commerce Backend API

A scalable and modular e-commerce backend API built with Go, following clean architecture principles and best practices.

## 🚀 Features

- RESTful API design
- Clean architecture implementation
- Modular project structure
- Database migrations support
- Logging system
- Configuration management
- Middleware support
- CLI and cron job capabilities

## 📁 Project Structure

```
├── cmd/                    # Application entry points
│   ├── cli/               # Command-line interface applications
│   ├── cronjob/           # Background job applications
│   └── server/            # HTTP server application
│       └── main.go        # Main server entry point
├── configs/               # Configuration files
├── docs/                  # Documentation files
├── global/                # Global variables and constants
├── internal/              # Private application code
│   ├── controller/        # HTTP request handlers
│   ├── initialize/        # Application initialization logic
│   ├── middlewares/       # HTTP middlewares
│   ├── models/            # Data models and structures
│   ├── repo/              # Repository layer (data access)
│   ├── router/            # HTTP route definitions
│   └── service/           # Business logic layer
├── migrations/            # Database migration files
│   └── 1_create_table_up.sql
├── pkg/                   # Public reusable packages
│   ├── logger/            # Logging utilities
│   ├── setting/           # Configuration management
│   └── utils/             # Common utilities
├── response/              # HTTP response structures
│   └── httpStatusCode.go  # HTTP status code definitions
├── scripts/               # Build and deployment scripts
├── tests/                 # Test files
├── third_party/           # External dependencies and integrations
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```

## 🏗️ Architecture

This project follows the **Clean Architecture** pattern with clear separation of concerns:

### Layers

1. **Controller Layer** (`internal/controller/`)

   - Handles HTTP requests and responses
   - Input validation and sanitization
   - Route handling logic

2. **Service Layer** (`internal/service/`)

   - Contains business logic
   - Orchestrates data flow between layers
   - Implements core application rules

3. **Repository Layer** (`internal/repo/`)

   - Data access abstraction
   - Database operations
   - External API integrations

4. **Model Layer** (`internal/models/`)
   - Data structures and entities
   - Domain models
   - Database schemas

### Key Directories

- **`cmd/`** - Application entry points for different executables
- **`internal/`** - Private application code (not importable by other projects)
- **`pkg/`** - Public reusable packages that can be imported by other projects
- **`global/`** - Global variables, constants, and shared configurations
- **`migrations/`** - Database schema migrations
- **`response/`** - Standardized HTTP response structures

## 🚦 Getting Started

### Prerequisites

- Go 1.21.3 or higher
- Database system (PostgreSQL, MySQL, etc.)
- Git

### Installation

1. Clone the repository:

```bash
git clone https://github.com/AlbertBui010/go-ecommerce-backend-api.git
cd go-ecommerce-backend-api
```

2. Install dependencies:

```bash
go mod tidy
```

3. Set up your environment configuration in the `configs/` directory

4. Run database migrations:

```bash
# Add your migration command here
```

5. Start the server:

```bash
go run cmd/server/main.go
```

## 🔧 Configuration

Configuration files are stored in the `configs/` directory. The application uses the `pkg/setting/` package for configuration management.

## 📊 Database

Database migrations are managed in the `migrations/` directory. Each migration file follows the naming convention:

- `{version}_{description}_up.sql` - Forward migration
- `{version}_{description}_down.sql` - Rollback migration

## 🧪 Testing

Run tests using:

```bash
go test ./...
```

Test files are organized in the `tests/` directory and follow Go testing conventions.

## 📝 API Documentation

API documentation is generated and stored in the `docs/` directory.

## 🔒 Middleware

The application supports various middleware for:

- Authentication and authorization
- Logging and monitoring
- Rate limiting
- CORS handling
- Request/response transformation

Middleware implementations are in `internal/middlewares/`.

## 🏃‍♂️ CLI Commands

The project includes CLI commands for various operations:

```bash
go run cmd/cli/main.go [command] [options]
```

## ⏰ Background Jobs

Cron jobs and background tasks are implemented in `cmd/cronjob/` for:

- Data synchronization
- Cleanup tasks
- Scheduled reports
- Email notifications

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

Project Link: [AlbertBui010/go-ecommerce-backend-api](https://github.com/AlbertBui010/go-ecommerce-backend-api)
