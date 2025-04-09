package client

import (
	"context"

	productpb "github.com/yourproject/services/product/gen"
	"google.golang.org/grpc"
)

type ProductClient interface {
	GetProducts(ctx context.Context, ids []string) ([]model.Product, float64, error)
}

type GRPCProductClient struct {
	conn *grpc.ClientConn
}

func NewProductClient(addr string) (*GRPCProductClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &GRPCProductClient{conn: conn}, nil
}

func (c *GRPCProductClient) GetProducts(ctx context.Context, ids []string) ([]model.Product, float64, error) {
	client := productpb.NewProductServiceClient(c.conn)

	resp, err := client.GetProducts(ctx, &productpb.ProductQuery{
		Ids: ids,
	})
	if err != nil {
		return nil, 0, err
	}

	var total float64
	products := make([]model.Product, len(resp.Products))

	for i, p := range resp.Products {
		products[i] = model.Product{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		}
		total += p.Price
	}

	return products, total, nil
}
