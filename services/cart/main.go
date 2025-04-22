package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"cart/internal/db"
	"cart/internal/handler"
	pb "cart/proto/gen/cart"

	productpb "cart/proto/gen/product"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Initialize database
	log.Println("Connecting to database...")
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	// Create gRPC server with health checks
	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	// Register services
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productServiceAddr := os.Getenv("PRODUCT_SERVICE_ADDR")
	if productServiceAddr == "" {
		productServiceAddr = "product-service:50052" // Default if not set
	}

	productConn, err := grpc.DialContext(ctx,
		productServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	// Create cart handler with dependencies
	cartHandler := handler.NewCartHandler(
		dbInstance,
		productpb.NewProductServiceClient(productConn),
	)
	pb.RegisterCartServiceServer(grpcServer, cartHandler)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	// Set health status
	healthServer.SetServingStatus(pb.CartService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// Start server
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Cart service listening on port 50053...")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down cart service...")
	healthServer.SetServingStatus(pb.CartService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_NOT_SERVING)
	grpcServer.GracefulStop()
}
