package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cec-project/delivery-app/services/order/config"
	"cec-project/delivery-app/services/order/internal/client"
	"cec-project/delivery-app/services/order/internal/handler"
	"cec-project/delivery-app/services/order/internal/repository"
	"cec-project/delivery-app/services/order/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize file storage
	store := storage.NewFileStore(cfg.DataDir)

	// Initialize product client
	productClient := client.NewProductClient(cfg.ProductServiceAddr)

	// Repository
	orderRepo := repository.NewOrderRepository(store, productClient)

	// GRPC Server
	grpcServer := grpc.NewServer()
	handler.RegisterOrderServiceServer(grpcServer, handler.NewOrderHandler(orderRepo))

	// Start server
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("Order service running on port %s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down order service...")
	grpcServer.GracefulStop()
}
