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
	grpcServer := grpc.NewServer()
	orderHandler := handler.NewOrderHandler()
}
