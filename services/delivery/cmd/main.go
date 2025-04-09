package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cec-project/delivery-app/services/delivery/config"
	"cec-project/delivery-app/services/delivery/internal/client"
	"cec-project/delivery-app/services/delivery/internal/handler"
	"cec-project/delivery-app/services/delivery/internal/repository"
	"cec-project/delivery-app/services/delivery/internal/storage"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize file storage
	store := storage.NewFileStore(cfg.DataDir)

	// Initialize order client
	orderClient := client.NewOrderClient(cfg.OrderServiceAddr)

	// Repository
	deliveryRepo := repository.NewDeliveryRepository(store, orderClient)

	// GRPC Server
	grpcServer := grpc.NewServer()
	handler.RegisterDeliveryServiceServer(grpcServer, handler.NewDeliveryHandler(deliveryRepo))

	// Start server
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("Delivery service running on port %s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Shutting down delivery service...")
	grpcServer.GracefulStop()
}
