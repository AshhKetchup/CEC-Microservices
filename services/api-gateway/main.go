package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "cec/service/gateway/proto/gen"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	// Register Auth Service
	authConn, _ := grpc.DialContext(
		ctx,
		"auth-service:50051", // Docker service name
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	authpb.RegisterAuthServiceHandlerClient(ctx, mux, authpb.NewAuthServiceClient(authConn))

	// productConn, _ := grpc.DialContext(
	// 	ctx,
	// 	"product-service:50052",
	// 	grpc. WithTransportCredentials(insecure.NewCredentials()),
	// )
	// productpb.RegisterProductServiceServer(ctx, mux, NewProductServiceClient(productConn))

	// Add middleware
	handler := cors(mux)

	// Start HTTP server
	http.ListenAndServe(":8080", handler)
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
