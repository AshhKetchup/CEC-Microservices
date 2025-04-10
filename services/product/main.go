package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cec/services/product/internal/db"
	"cec/services/product/internal/handler"
	pb "cec/services/product/proto/gen"

	"google.golang.org/grpc"
)

func main() {
	// Initialize database
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	// Create gRPC server
	grpcServer := grpc.NewServer()
	productHandler := handler.NewProductHandler(dbInstance)
	pb.RegisterProductServiceServer(grpcServer, productHandler)

	// Start server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

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
	grpcServer.GracefulStop()
}
