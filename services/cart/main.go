package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cart/internal/db"
	"cart/internal/handler"
	pb "cart/proto/gen"

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
	cartHandler := handler.NewCartHandler(dbInstance)
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
