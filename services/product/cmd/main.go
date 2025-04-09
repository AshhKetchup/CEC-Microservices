package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cec/services/product/internals"
	pb "cec/services/product/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
}

func main() {
	// Initialize database connection using shared package
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbInstance.Close()

	// Create products table if not exists
	_, err = dbInstance.Exec(`CREATE TABLE IF NOT EXISTS products (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &productServer{})

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

// CreateProduct implementation
func (s *productServer) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	dbInstance := db.GetDB()

	result, err := dbInstance.ExecContext(ctx,
		"INSERT INTO products (id, name, description, price) VALUES (UUID(), ?, ?, ?)",
		req.Name, req.Description, req.Price,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product ID: %v", err)
	}

	var product pb.ProductResponse
	err = dbInstance.QueryRowContext(ctx,
		"SELECT id, name, description, price FROM products WHERE id = ?",
		lastInsertId,
	).Scan(&product.Id, &product.Name, &product.Description, &product.Price)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve created product: %v", err)
	}

	return &product, nil
}

// GetProduct implementation
func (s *productServer) GetProduct(ctx context.Context, req *pb.ProductID) (*pb.ProductResponse, error) {
	dbInstance := db.GetDB()

	var product pb.ProductResponse
	err := dbInstance.QueryRowContext(ctx,
		"SELECT id, name, description, price FROM products WHERE id = ?",
		req.Id,
	).Scan(&product.Id, &product.Name, &product.Description, &product.Price)

	switch {
	case err == db.ErrNoRows:
		return nil, status.Error(codes.NotFound, "product not found")
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to get product: %v", err)
	}

	return &product, nil
}

// ListProducts implementation
func (s *productServer) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductListResponse, error) {
	dbInstance := db.GetDB()

	rows, err := dbInstance.QueryContext(ctx,
		"SELECT id, name, description, price FROM products",
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}
	defer rows.Close()

	var products []*pb.ProductResponse
	for rows.Next() {
		var p pb.ProductResponse
		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to scan product: %v", err)
		}
		products = append(products, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "error after scanning rows: %v", err)
	}

	return &pb.ProductListResponse{Products: products}, nil
}
