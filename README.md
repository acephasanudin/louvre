# Louvre Boilerplate Template

A comprehensive Go microservice boilerplate template based on clean architecture principles with multiple transport layers, database integration, and production-ready features.

## Features

- **Clean Architecture**: Hexagonal architecture with clear separation of concerns
- **Multiple Transport Layers**: REST API, gRPC, Cron Jobs, and Kafka message handling
- **Database Integration**: PostgreSQL with GORM ORM and migrations
- **Authentication & Authorization**: Keycloak integration
- **Caching**: Redis integration
- **Tracing**: OpenTelemetry with Jaeger
- **Configuration Management**: YAML-based configuration with environment support
- **Dependency Injection**: Google Wire for compile-time DI
- **API Documentation**: Swagger/OpenAPI with Fiber framework
- **Testing**: Comprehensive test setup with coverage
- **Docker Support**: Multi-environment Docker configurations

## Requirements

- Go (version 1.23 or higher)
- PostgreSQL
- Keycloak
- Redis
- Kafka
- Docker (optional, for running services)

## Project Structure

```
├── cmd/                          # Command-line interface
│   ├── cmd.go                   # Root command setup
│   └── services/
│       └── start.go             # Service startup logic
├── app/                         # Application layer
│   ├── adapter/                 # Dependency injection adapters
│   ├── service/                 # Service implementations
│   └── port/                    # Port interfaces
├── internal/                    # Internal application code
│   ├── config/                  # Configuration management
│   ├── domain/                  # Domain logic
│   │   └── example/             # Example domain
│   │       ├── model/           # Data models
│   │       ├── repository/      # Repository interfaces
│   │       └── usecase/         # Business logic
│   ├── handler/                 # HTTP/gRPC handlers
│   ├── middleware/              # Middleware implementations
│   ├── model/                   # Shared models
│   └── service/                 # Service layer
├── pkg/                         # Reusable packages
│   ├── cache/                   # Cache implementations
│   ├── database/                # Database utilities
│   ├── env/                     # Environment management
│   ├── kafka/                   # Kafka client
│   ├── keycloak/                # Keycloak integration
│   ├── logger/                  # Logging utilities
│   └── tracing/                 # OpenTelemetry tracing
├── config/                      # Configuration files
│   └── file/                    # Environment-specific configs
├── proto/                       # Protocol buffer definitions
├── deployment/                  # Kubernetes deployment files
├── build/                       # Build artifacts
└── scripts/                     # Utility scripts
```

## Getting Started

1. **Clone the template:**
   ```bash
   git clone <repository-url>
   cd go-boilerplate-template
   ```

2. **Update module name:**
   - Replace `example/service` in `go.mod` with your actual module name
   - Update all import statements throughout the codebase

3. **Install dependencies:**
   ```bash
   go mod download
   ```

4. **Configure the application:**
   - Copy and modify configuration files in `config/file/`
   - Update database connections, API keys, and other settings

5. **Run the application:**
   ```bash
   # Run all services
   go run main.go start

   # Run specific services
   go run main.go start --config config/file/config.local.yaml
   ```

## Available Commands

### Development
```bash
# Start all services (REST, gRPC, Cron, Kafka)
make dev

# Generate API documentation
make doc

# Run tests with coverage
make test
```

### Code Generation
```bash
# Generate dependency injection code
make wire-gen

# Generate protobuf files
make buf-gen
```

### Database Migrations
```bash
# Run database migrations
go run main.go migration up

# Create new migration
go run main.go migration create migration_name
```

## Configuration

The application uses YAML configuration files located in `config/file/`:

- `config.local.yaml` - Local development
- `config.dev.yaml` - Development environment
- `config.docker.yaml` - Docker environment

### Configuration Sections

- **serviceName**: Name of your service
- **transport**: Transport layer configurations (REST, gRPC, Kafka, Cron)
- **datasource**: External service configurations (PostgreSQL, Keycloak, APIs)
- **tracer**: OpenTelemetry tracing configuration
- **filestorage**: File storage settings

## Adding New Features

### 1. Create New Domain

```bash
# Create domain structure
mkdir -p internal/domain/newfeature/{model,repository,usecase}
```

### 2. Define Model

Create your model in `internal/domain/newfeature/model/`:

```go
type NewFeature struct {
    ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
    Name      string    `gorm:"not null"`
    // ... other fields
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### 3. Define Repository Interface

Create repository interface in `internal/domain/newfeature/repository/`:

```go
type NewFeatureRepository interface {
    Create(ctx context.Context, feature *model.NewFeature) error
    GetByID(ctx context.Context, id uuid.UUID) (*model.NewFeature, error)
    // ... other methods
}
```

### 4. Implement Use Case

Create use case in `internal/domain/newfeature/usecase/`:

```go
type NewFeatureUseCase interface {
    CreateFeature(ctx context.Context, req *CreateFeatureRequest) (*model.NewFeature, error)
    // ... other methods
}
```

### 5. Add to Wire Configuration

Update `app/adapter/wire_set.go` to include your new providers.

## Environment Variables

Key environment variables to consider:

- `CONFIG_PATH`: Path to configuration directory
- `ENVIRONMENT`: Environment name (development, staging, production)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)

## Docker Support

The template includes Docker configurations for different environments:

```bash
# Build Docker image
docker build -t your-service-name .

# Run with Docker Compose
docker-compose up -d
```

## Monitoring & Observability

- **Logging**: Structured logging with zerolog
- **Tracing**: OpenTelemetry with Jaeger integration
- **Metrics**: Built-in metrics collection (extend as needed)
- **Health Checks**: Health check endpoints for monitoring

## Testing

The template includes a comprehensive testing setup:

```bash
# Run all tests
go test ./...

# Run tests with coverage
make test

# Run specific package tests
go test ./internal/domain/example/usecase/
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For questions and support, please create an issue in the repository.

---

**Note**: This is a boilerplate template. Remember to:
- Replace placeholder text and module names
- Update configurations for your specific needs
- Implement actual business logic
- Add proper error handling
- Configure security settings appropriately
- Set up CI/CD pipelines
- Add comprehensive documentation
