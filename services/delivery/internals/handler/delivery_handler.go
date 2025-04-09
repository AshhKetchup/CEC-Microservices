package handler

import (
	"context"
	"time"

	pb "cec-project/delivery-app/services/delivery/gen"
	"cec-project/delivery-app/services/delivery/internal/model"
	"cec-project/delivery-app/services/delivery/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryHandler struct {
	repo repository.DeliveryRepository
	pb.UnimplementedDeliveryServiceServer
}

func NewDeliveryHandler(repo repository.DeliveryRepository) *DeliveryHandler {
	return &DeliveryHandler{repo: repo}
}

func (h *DeliveryHandler) CreateDelivery(ctx context.Context, req *pb.DeliveryRequest) (*pb.DeliveryResponse, error) {
	delivery, err := h.repo.CreateDelivery(ctx, model.DeliveryRequest{
		OrderID: req.OrderId,
		Address: req.Address,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeliveryResponse{
		Delivery: &pb.Delivery{
			Id:                delivery.ID,
			Order:             convertToPBOrder(delivery.Order),
			Address:           delivery.Address,
			Status:            pb.DeliveryStatus(delivery.Status),
			EstimatedDelivery: convertToPBTimestamp(delivery.EstimatedDelivery),
			CreatedAt:         convertToPBTimestamp(delivery.CreatedAt),
		},
		Base: &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

// Implement other methods (GetDelivery, UpdateDeliveryStatus)
