// cart_handler.go
package handler

import (
	"context"
	"database/sql"
	"time"

	pb "cart/proto/gen/cart"
	productpb "cart/proto/gen/product"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartHandler struct {
	db            *sql.DB
	productClient productpb.ProductServiceClient // Add product client
	pb.UnimplementedCartServiceServer
}

func NewCartHandler(db *sql.DB, productClient productpb.ProductServiceClient) *CartHandler {
	return &CartHandler{
		db:            db,
		productClient: productClient, // Initialize product client
	}
}

func (h *CartHandler) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.BaseResponse, error) {
	// Check if product exists
	//productServiceAddr := os.Getenv("PRODUCT_SERVICE_ADDR")

	_, err := h.productClient.GetProduct(ctx, &productpb.ProductID{
		Id: req.ProductId,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "product service error: %v", err)
	}

	// Product exists, proceed with adding to cart
	_, err = h.db.ExecContext(ctx,
		`INSERT INTO cart_items (user_id, product_id, quantity) 
        VALUES (?, ?, ?)
        ON DUPLICATE KEY UPDATE quantity = quantity + ?`,
		req.UserId, req.ProductId, req.Quantity, req.Quantity,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart: %v", err)
	}

	return &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Item added to cart successfully",
	}, nil
}

func (h *CartHandler) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	// Get cart total with COALESCE
	var total float64
	err = tx.QueryRowContext(ctx,
		`SELECT COALESCE(SUM(p.price * ci.quantity), 0) 
        FROM cart_items ci
        JOIN products p ON ci.product_id = p.id
        WHERE ci.user_id = ?`,
		req.UserId,
	).Scan(&total)

	// Removed redundant error check

	// Validate cart not empty
	if total <= 0 {
		return nil, status.Error(
			codes.FailedPrecondition,
			"cannot process payment for empty cart",
		)
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to calculate total: %v", err)
	}

	// Create payment record
	txID := uuid.New().String()
	_, err = tx.ExecContext(ctx,
		`INSERT INTO payments (transaction_id, user_id, amount, currency, status) 
		VALUES (?, ?, ?, ?, 'COMPLETED')`,
		txID, req.UserId, total, req.Currency,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "payment processing failed: %v", err)
	}

	// Clear cart
	_, err = tx.ExecContext(ctx,
		"DELETE FROM cart_items WHERE user_id = ?",
		req.UserId,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to clear cart: %v", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, status.Errorf(codes.Internal, "transaction commit failed: %v", err)
	}

	return &pb.PaymentResponse{
		TransactionId: txID,
		PaymentTime: &pb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   int32(time.Now().Nanosecond()),
		},
		Base: &pb.BaseResponse{
			Code:    int32(codes.OK),
			Message: "Payment processed successfully",
		},
	}, nil
}
