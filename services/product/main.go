package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"product/internal/db"
	"product/internal/handler"
	pb "product/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Initialize database with retries
	log.Println("Waiting for MySQL to become available...")
	//time.Sleep(20 * time.Second)

	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	log.Println("Database done now in main func.")
	// Create gRPC server with health checks
	grpcServer := grpc.NewServer()
	healthServer := health.NewServer()

	// Register services
	productHandler := handler.NewProductHandler(dbInstance)
	pb.RegisterProductServiceServer(grpcServer, productHandler)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	// Set health status
	healthServer.SetServingStatus(pb.ProductService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	print("pb func done")
	// Start server on all interfaces
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening on port 50052...")
	go func() {
		log.Println("Product service running on port 50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down product service...")
	healthServer.SetServingStatus(pb.ProductService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_NOT_SERVING)
	grpcServer.GracefulStop()
}
