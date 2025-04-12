package main

import (
	"auth/internal/db"
	"auth/internal/handler"
	pb "auth/proto/gen"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	// Create gRPC server
	grpcServer := grpc.NewServer()
	authHandler := handler.NewAuthHandler(dbInstance)
	pb.RegisterAuthServiceServer(grpcServer, authHandler)

	reflection.Register(grpcServer)

	// Start server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Println("Auth service running on port 50052")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down auth service...")
	grpcServer.GracefulStop()
}
