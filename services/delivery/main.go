package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"delivery/internal/db"
	"delivery/internal/handler"
	pb "delivery/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Initialize database
	log.Println("Initializing delivery database...")
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	// Create gRPC server with health checks
	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	// Register services
	deliveryHandler := handler.NewDeliveryHandler(dbInstance)
	pb.RegisterDeliveryServiceServer(grpcServer, deliveryHandler)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	// Set health status
	healthServer.SetServingStatus(pb.DeliveryService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// Start server
	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Delivery service listening on port 50054...")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down delivery service...")
	healthServer.SetServingStatus(pb.DeliveryService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_NOT_SERVING)
	grpcServer.GracefulStop()
}
