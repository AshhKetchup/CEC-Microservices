// delivery_handler.go
package handler

import (
	"context"
	"database/sql"
	"time"

	pb "delivery/proto/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeliveryHandler struct {
	db *sql.DB
	pb.UnimplementedDeliveryServiceServer
}

func NewDeliveryHandler(db *sql.DB) *DeliveryHandler {
	return &DeliveryHandler{db: db}
}

func protoToTime(ts *pb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
}

func timeToProto(t time.Time) *pb.Timestamp {
	return &pb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func (h *DeliveryHandler) ScheduleDelivery(ctx context.Context, req *pb.DeliveryRequest) (*pb.DeliveryResponse, error) {
	deliveryID := uuid.New().String()

	// Use custom conversion
	scheduleTime := protoToTime(req.ScheduleTime)

	_, err := h.db.ExecContext(ctx,
		`INSERT INTO deliveries (id, order_id, schedule_time, address, status)
		VALUES (?, ?, ?, ?, ?)`,
		deliveryID, req.OrderId, scheduleTime,
		req.DeliveryAddress, pb.DeliveryStatus_SCHEDULED.String(),
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to schedule delivery: %v", err)
	}

	return &pb.DeliveryResponse{
		DeliveryId: deliveryID,
		OrderId:    req.OrderId,
		Status:     pb.DeliveryStatus_SCHEDULED,
		ScheduledTime: &pb.Timestamp{
			Seconds: req.ScheduleTime.Seconds,
			Nanos:   req.ScheduleTime.Nanos,
		},
		Base: &pb.BaseResponse{
			Code:    int32(codes.OK),
			Message: "Delivery scheduled successfully",
		},
	}, nil
}

func (h *DeliveryHandler) UpdateDeliveryStatus(ctx context.Context, req *pb.StatusUpdate) (*pb.DeliveryResponse, error) {
	result, err := h.db.ExecContext(ctx,
		`UPDATE deliveries SET status = ?, updated_at = NOW() 
		WHERE id = ?`,
		req.Status.String(), req.DeliveryId,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update status: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "delivery not found")
	}

	return h.getDelivery(ctx, req.DeliveryId)
}

func (h *DeliveryHandler) TrackDelivery(ctx context.Context, req *pb.DeliveryQuery) (*pb.DeliveryResponse, error) {
	return h.getDelivery(ctx, req.DeliveryId)
}

func (h *DeliveryHandler) ConfirmDelivery(ctx context.Context, req *pb.DeliveryConfirmation) (*pb.BaseResponse, error) {
	_, err := h.db.ExecContext(ctx,
		`UPDATE deliveries 
		SET status = 'DELIVERED', 
			signature = ?,
			actual_delivery_time = NOW()
		WHERE id = ?`,
		req.Signature, req.DeliveryId,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to confirm delivery: %v", err)
	}

	return &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Delivery confirmed successfully",
	}, nil
}

func (h *DeliveryHandler) getDelivery(ctx context.Context, deliveryID string) (*pb.DeliveryResponse, error) {
	var delivery pb.DeliveryResponse
	var statusStr string
	var scheduledTime time.Time
	var actualDeliveryTime sql.NullTime

	err := h.db.QueryRowContext(ctx,
		`SELECT id, order_id, status, schedule_time, actual_delivery_time
		FROM deliveries WHERE id = ?`,
		deliveryID,
	).Scan(
		&delivery.DeliveryId,
		&delivery.OrderId,
		&statusStr,
		&scheduledTime,
		&actualDeliveryTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "delivery not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get delivery: %v", err)
	}

	delivery.Status = pb.DeliveryStatus(pb.DeliveryStatus_value[statusStr])
	delivery.ScheduledTime = timeToProto(scheduledTime)

	if actualDeliveryTime.Valid {
		delivery.ActualDeliveryTime = timeToProto(actualDeliveryTime.Time)
	}

	delivery.Base = &pb.BaseResponse{
		Code:    int32(codes.OK),
		Message: "Delivery details retrieved",
	}

	return &delivery, nil
}
