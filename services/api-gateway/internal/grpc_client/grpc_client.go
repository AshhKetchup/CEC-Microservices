package grpc_client

import (
	"context"

	authpb "cec-project/delivery-app/api/gen/auth"
	deliverypb "cec-project/delivery-app/api/gen/delivery"
	productpb "cec-project/delivery-app/api/gen/product"
	orderpb "cec-project/delivery-app/gen/order"

	"google.golang.org/grpc"
)

type Clients struct {
	AuthClient     authpb.AuthServiceClient
	ProductClient  productpb.ProductServiceClient
	OrderClient    orderpb.OrderServiceClient
	DeliveryClient deliverypb.DeliveryServiceClient
}

func NewClients(ctx context.Context) (*Clients, error) {
	authConn, err := grpc.Dial("auth-service:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	productConn, err := grpc.Dial("product-service:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	orderConn, err := grpc.Dial("order-service:50053", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	deliveryConn, err := grpc.Dial("delivery-service:50054", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Clients{
		AuthClient:     authpb.NewAuthServiceClient(authConn),
		ProductClient:  productpb.NewProductServiceClient(productConn),
		OrderClient:    orderpb.NewOrderServiceClient(orderConn),
		DeliveryClient: deliverypb.NewDeliveryServiceClient(deliveryConn),
	}, nil
}
