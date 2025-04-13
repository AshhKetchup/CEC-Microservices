package handler

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "product/proto/gen/"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandler struct {
	db *sql.DB
	pb.UnimplementedProductServiceServer
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	//var createdAt time.Time
	id := uuid.New().String()

	// Manually provide the ID since MySQL doesn't auto-generate UUIDs
	_, err := h.db.ExecContext(ctx,
		`INSERT INTO products (id, name, description, price, created_at) 
         VALUES (?, ?, ?, ?, ?)`,
		id, req.Name, req.Description, req.Price, time.Now(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	// Return constructed ProductResponse
	return &pb.ProductResponse{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt: &pb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   int32(time.Now().Nanosecond()),
		},
		Base: &pb.BaseResponse{
			Code:    int32(codes.OK),
			Message: "Product created successfully",
		},
	}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.ProductID) (*pb.ProductResponse, error) {
	var product pb.ProductResponse
	var createdAt time.Time

	err := h.db.QueryRowContext(ctx,
		`SELECT id, name, description, price, created_at 
         FROM products 
         WHERE id = ?`,
		req.Id,
	).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &createdAt)

	switch {
	case err == sql.ErrNoRows:
		return nil, status.Error(codes.NotFound, "product not found")
	case err != nil:
		return nil, status.Errorf(codes.Internal,
			"failed to get product: %v", err)
	}

	product.CreatedAt = &pb.Timestamp{
		Seconds: createdAt.Unix(),
		Nanos:   int32(createdAt.Nanosecond()),
	}

	product.Base = &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Product retrieved successfully",
	}

	return &product, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.ProductQuery) (*pb.ProductListResponse, error) {
	response := &pb.ProductListResponse{
		Products: []*pb.ProductResponse{},
		Base: &pb.BaseResponse{
			Code: int32(codes.OK),
		},
	}

	// Calculate offset
	offset := (req.Page - 1) * req.Limit

	// Get products
	rows, err := h.db.QueryContext(ctx,
		`SELECT id, name, description, price, created_at 
         FROM products 
         LIMIT ? OFFSET ?`,
		req.Limit, offset,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to list products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product pb.ProductResponse
		var createdAt time.Time

		if err := rows.Scan(
			&product.Id, &product.Name,
			&product.Description, &product.Price, &createdAt,
		); err != nil {
			return nil, status.Errorf(codes.Internal,
				"failed to scan product: %v", err)
		}

		product.CreatedAt = &pb.Timestamp{
			Seconds: createdAt.Unix(),
			Nanos:   int32(createdAt.Nanosecond()),
		}

		response.Products = append(response.Products, &product)
	}

	// Get total count
	var total int32
	err = h.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM products",
	).Scan(&total)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to get total products: %v", err)
	}

	response.Total = total
	response.Base.Message = fmt.Sprintf("Found %d products", total)

	return response, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.ProductUpdate) (*pb.ProductResponse, error) {
	var product pb.ProductResponse
	var createdAt time.Time

	// Check if product exists
	err := h.db.QueryRowContext(ctx,
		"SELECT id FROM products WHERE id = ?",
		req.Id,
	).Scan(&product.Id)

	switch {
	case err == sql.ErrNoRows:
		return nil, status.Error(codes.NotFound, "product not found")
	case err != nil:
		return nil, status.Errorf(codes.Internal,
			"failed to verify product: %v", err)
	}

	// Update product
	result, err := h.db.ExecContext(ctx,
		`UPDATE products 
         SET name = ?, description = ?, price = ? 
         WHERE id = ?`,
		req.Name, req.Description, req.Price, req.Id,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to update product: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	// Get updated product
	err = h.db.QueryRowContext(ctx,
		`SELECT id, name, description, price, created_at 
         FROM products 
         WHERE id = ?`,
		req.Id,
	).Scan(&product.Id, &product.Name, &product.Description,
		&product.Price, &createdAt)

	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to get updated product: %v", err)
	}

	product.CreatedAt = &pb.Timestamp{
		Seconds: createdAt.Unix(),
		Nanos:   int32(createdAt.Nanosecond()),
	}

	product.Base = &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Product updated successfully",
	}

	return &product, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.ProductID) (*pb.BaseResponse, error) {
	// Check if product exists
	var id string
	err := h.db.QueryRowContext(ctx,
		"SELECT id FROM products WHERE id = ?",
		req.Id,
	).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return nil, status.Error(codes.NotFound, "product not found")
	case err != nil:
		return nil, status.Errorf(codes.Internal,
			"failed to verify product: %v", err)
	}

	// Delete product
	_, err = h.db.ExecContext(ctx,
		"DELETE FROM products WHERE id = ?",
		req.Id,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to delete product: %v", err)
	}

	return &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Product deleted successfully",
	}, nil
}
