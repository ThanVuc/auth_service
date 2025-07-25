# Internal and Script Directories

This project uses a modular structure to separate core logic and utility scripts. Below is a focused overview of the `internal/` and `script/` directories.

## internal/
Contains the main application logic, organized by responsibility:

- **amqp/**
  - `consumers/`: AMQP consumer implementations (e.g., resource API consumer, consumer factory)
  - `producers/`: (empty) Placeholder for AMQP producer logic
- **controller/**: HTTP/gRPC controllers for authentication, permissions, roles, and tokens
- **database/**: Database models and initialization logic
- **grpc/**: gRPC service definitions and generated code for auth, permission, role, and token
- **helper/**: Helper functions (e.g., resource collection)
- **initialize/**: Initialization logic for consumers, logger, PostgreSQL (GORM and SQLC), RabbitMQ, Redis, gRPC server, and configuration loading
- **middlewares/**: Middleware components (empty or for future use)
- **models/**: Data models (e.g., authentication log model)
- **repos/**: Repository layer for data access (auth, permission, role, token)
- **services/**: Business logic/services (auth, permission, role, token)
- **wire/**: Dependency injection setup using wire (auth, permission, role, token, and generated code)

## script/
Utility scripts to automate common development tasks:

- `init_proto.sh`: Initializes protocol buffers for gRPC services
- `wire-gen.sh`: Generates dependency injection code using wire

## internal/initialize/
This directory is responsible for initializing and wiring up the application's core infrastructure and services. Each file has a specific role in preparing the environment before the main business logic runs:

- **init.consumer.go**: Initializes and starts all AMQP consumers using a consumer factory. Ensures background consumers are running for message processing.
- **init.logger.go**: Sets up the global logger instance using configuration values. All logging throughout the app uses this logger.
- **init.postgreSql.gorm.go**: (Obsolete) Previously used to initialize PostgreSQL with GORM. Now replaced by SQLC and pgx for database access.
- **init.postgreSql.sqlc.go**: Initializes the PostgreSQL connection pool using pgx and stores it globally for use by repositories and services.
- **init.rabbitMQ.go**: Establishes a connection to RabbitMQ and sets up a shared channel for message publishing and consumption.
- **init.redis.go**: Initializes the Redis client and stores it globally for caching and other Redis operations.
- **init.server.grpc.go**: Sets up and runs the gRPC server. Wires in all controllers (Auth, Permission, Role, Token) using dependency injection, registers them with the gRPC server, and manages server lifecycle and graceful shutdown.
- **load.config.go**: Loads application configuration from YAML files using Viper and stores it in a global config variable.
- **run.go**: The main entry point for application startup. Loads configuration, initializes all dependencies (logger, database, Redis, RabbitMQ), starts the gRPC server, launches consumers, and manages graceful shutdown on OS signals.

### How it all connects
- `run.go` orchestrates the startup process, calling each initialization function in order.
- `init.server.grpc.go` is called by `run.go` to start the gRPC server, which registers controllers that handle incoming requests.
- Controllers (from `internal/controller/`) are injected into the gRPC server and serve as endpoints for business logic, using services and repositories.
- All infrastructure (database, Redis, RabbitMQ, logger) is initialized before the server starts, ensuring the application is ready to handle requests and background jobs.

---

For more details on each component, see the respective subdirectory or script file.
