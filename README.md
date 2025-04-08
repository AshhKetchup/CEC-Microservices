# Microservices
> Building and Deploying a Microservices using Docker/Kubernetes Assignment under cloud and edge computing course (6th sem)

### Directory Structure
``` text
delivery-app/
├── api/                         # All proto definitions and generated code
│   ├── proto/
│   │   ├── auth.proto
│   │   ├── product.proto
│   │   ├── order.proto
│   │   ├── delivery.proto
│   │   └── common/
│   │       └── types.proto
│   └── gen/                     # Generated Go + gRPC + gRPC-Gateway code
│
├── pkg/                         # Shared reusable libraries
│   ├── logger/
│   ├── db/
│   ├── middleware/
│   └── auth/
│
├── services/
│   ├── auth/
│   │   ├── cmd/                 # Main entrypoint
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler.go       # Business logic
│   │   │   ├── server.go        # gRPC server
│   │   │   ├── model.go         # User model
│   │   │   └── db.go            # DB connection
│   │   ├── config/
│   │   │   └── config.yaml
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   └── go.sum
│
│   ├── product/
│   │   └── (same structure as auth/)
│
│   ├── order/
│   │   └── (same structure as auth/)
│
│   ├── delivery/
│   │   └── (same structure as auth/)
│
│   └── gateway/
│       ├── cmd/
│       │   └── main.go
│       ├── internal/
│       │   ├── middleware/      # JWT, logging, etc.
│       │   └── router.go        # Register gRPC-Gateway handlers
│       ├── config/
│       │   └── config.yaml
│       ├── Dockerfile
│       ├── go.mod
│       └── go.sum
│
├── deployments/
│   ├── docker-compose.yaml      # Run all services and DBs
│   └── k8s/                     # Optional: Kubernetes manifests
│
├── Makefile                     # Dev tasks
└── README.md

### Each service in detail
services/auth/
├── cmd/
│   └── main.go                  # Initializes server, loads config, starts gRPC
├── internal/
│   ├── handler.go               # Business logic (signup, login)
│   ├── server.go                # Implements gRPC server
│   ├── model.go                 # User structs
│   └── db.go                    # DB connection setup (Postgres/Mongo)
├── config/
│   └── config.yaml              # Env vars, DB creds, ports
├── Dockerfile                   # Containerize service
├── go.mod / go.sum              # Go modules


services/gateway/
├── cmd/
│   └── main.go                  # Starts HTTP server (gRPC-Gateway)
├── internal/
│   ├── router.go                # Register all gRPC handlers (auth, order, etc.)
│   ├── middleware/
│   │   ├── auth.go              # JWT auth middleware
│   │   └── logging.go
│   └── grpc_client/             # gRPC clients to talk to services if needed
├── config/
│   └── config.yaml
├── Dockerfile
├── go.mod / go.sum
```

📁 api/

    Purpose: Contains all .proto files and the generated gRPC + gRPC-Gateway Go code for communication.

Subdirectories:

    proto/ – Define all your gRPC service contracts.

        auth.proto, product.proto, etc.

        common/types.proto – Shared messages (e.g., timestamps, error responses)

    gen/ – Auto-generated Go code using protoc, split into:

        gRPC Go code

        gRPC-Gateway HTTP handlers



