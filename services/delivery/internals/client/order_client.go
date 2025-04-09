package client

import (
	"context"

	orderpb "cec-project/delivery-app/services/order/gen"
	"google.golang.org/grpc"
)

type OrderClient interface {
	GetOrder(ctx context.Context, orderID string) (*orderpb.OrderResponse, error)
}

type GRPCOrderClient struct {
	conn   *grpc.ClientConn
	client orderpb.OrderServiceClient
}

func NewOrderClient(addr string) (*GRPCOrderClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCOrderClient{
		conn:   conn,
		client: orderpb.NewOrderServiceClient(conn),
	}, nil
}

func (c *GRPCOrderClient) GetOrder(ctx context.Context, orderID string) (*orderpb.OrderResponse, error) {
	return c.client.GetOrder(ctx, &orderpb.OrderID{Id: orderID})
}

func (c *GRPCOrderClient) Close() error {
	return c.conn.Close()
}
