package handler

import (
	"context"

	"github.com/yourproject/services/order/internal/model"
	"github.com/yourproject/services/order/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/yourproject/services/order/gen"
)

type OrderHandler struct {
	repo repository.OrderRepository
	pb.UnimplementedOrderServiceServer
}

func NewOrderHandler(repo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{repo: repo}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	order, err := h.repo.CreateOrder(ctx, model.OrderRequest{
		UserID:     req.UserId,
		ProductIDs: req.ProductIds,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.OrderResponse{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     convertToPBProducts(order.Items),
		Total:     order.Total,
		Status:    pb.OrderStatus(order.Status),
		CreatedAt: convertToPBTimestamp(order.CreatedAt),
		Base:      &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.OrderResponse, error) {
	order, err := h.repo.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &pb.OrderResponse{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     convertToPBProducts(order.Items),
		Total:     order.Total,
		Status:    pb.OrderStatus(order.Status),
		CreatedAt: convertToPBTimestamp(order.CreatedAt),
		Base:      &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, req *pb.OrderStatusUpdate) (*pb.OrderResponse, error) {
	order, err := h.repo.UpdateOrderStatus(ctx, req.OrderId, model.OrderStatus(req.Status))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.OrderResponse{
		Id:     order.ID,
		Status: pb.OrderStatus(order.Status),
		Base:   &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

// Helper functions for proto conversions
func convertToPBProducts(items []model.Product) []*pb.ProductResponse {
	// Implementation
}
