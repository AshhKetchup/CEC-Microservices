package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "api-gateway/gen/auth"
	cartpb "api-gateway/gen/cart"
	deliverypb "api-gateway/gen/delivery"
	productpb "api-gateway/gen/product"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gRPC-Gateway mux
	gwMux := runtime.NewServeMux()

	// Common dial options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Get service addresses from environment variables
	authAddr := os.Getenv("AUTH_SERVICE_ADDR")
	productAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	cartAddr := os.Getenv("CART_SERVICE_ADDR")
	deliveryAddr := os.Getenv("DELIVERY_SERVICE_ADDR")

	// Register all gRPC service handlers

	if err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwMux, authAddr, opts); err != nil {
		log.Fatalf("Failed to register AuthService handler: %v", err)
	}

	if err := productpb.RegisterProductServiceHandlerFromEndpoint(ctx, gwMux, productAddr, opts); err != nil {
		log.Fatalf("Failed to register ProductService handler: %v", err)
	}

	if err := cartpb.RegisterCartServiceHandlerFromEndpoint(ctx, gwMux, cartAddr, opts); err != nil {
		log.Fatalf("Failed to register CartService handler: %v", err)
	}

	if err := deliverypb.RegisterDeliveryServiceHandlerFromEndpoint(ctx, gwMux, deliveryAddr, opts); err != nil {
		log.Fatalf("Failed to register DeliveryService handler: %v", err)
	}

	// Optional REST endpoints (e.g. health check)
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Mount gRPC-Gateway mux
	httpMux.Handle("/", gwMux)

	log.Println("API Gateway is listening on :8080")
	if err := http.ListenAndServe(":8080", httpMux); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
