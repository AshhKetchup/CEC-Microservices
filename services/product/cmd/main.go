package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cec-project/delivery-app/services/product/config"
	"cec-project/delivery-app/services/product/internal/handler"
	"cec-project/delivery-app/services/product/internal/repository"
	"cec-project/delivery-app/services/product/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize file storage
	store := storage.NewFileStore(cfg.DataDir)

	// Repository
	productRepo := repository.NewProductRepository(store)

	// GRPC Server
	grpcServer := grpc.NewServer()
	handler.RegisterProductServiceServer(grpcServer, handler.NewProductHandler(productRepo))

	// Start server
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("Product service running on port %s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down product service...")
	grpcServer.GracefulStop()
}
