package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Update the import path below to match the generated package location.
	productpb "cec/service/api-gateway/gen/product"
)

func main() {
	// Create a root context.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a new gRPC-Gateway mux.
	gwMux := runtime.NewServeMux()

	// Dial options for connecting insecurely.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Register the ProductService endpoint.
	// The endpoint here ("product-service:50052") should match the target
	// gRPC server. In a Docker environment, this would be the name of the service.
	err := productpb.RegisterProductServiceHandlerFromEndpoint(
		ctx,
		gwMux,
		"product-service:50052",
		opts,
	)
	if err != nil {
		log.Fatalf("Failed to register ProductService handler: %v", err)
	}

	// You can also add additional REST endpoints.
	// For example, a simple health check.
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	// Mount the gRPC-Gateway mux on the root.
	httpMux.Handle("/", gwMux)

	// Start the HTTP server (API Gateway).
	log.Println("API Gateway is listening on :8080")
	if err := http.ListenAndServe(":8080", httpMux); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
