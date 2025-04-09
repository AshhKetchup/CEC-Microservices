package main

import (
	"context"
	"net/http"

	"cec-project/delivery-app/api-gateway/internal/middleware"
	"cec-project/delivery-app/api-gateway/internal/router"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize gRPC clients
	clients, err := grpc_client.NewClients(ctx)
	if err != nil {
		logger.Fatal("failed to create grpc clients", zap.Error(err))
	}

	// Create router with middleware chain
	r := router.NewRouter(clients)
	r.Use(middleware.Logging(logger))
	r.Use(middleware.Auth(clients.AuthClient))

	// Start HTTP server
	port := ":8080"
	logger.Info("starting gateway server", zap.String("port", port))
	if err := http.ListenAndServe(port, r); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
