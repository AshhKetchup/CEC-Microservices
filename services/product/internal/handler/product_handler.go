package handler

import (
	"context"

	"github.com/yourproject/services/product/internal/model"
	"github.com/yourproject/services/product/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "cec-project/delivery-app/services/product/gen"
)

type ProductHandler struct {
	repo repository.ProductRepository
	pb.UnimplementedProductServiceServer
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	product, err := h.repo.CreateProduct(ctx, model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   convertToPBTimestamp(product.CreatedAt),
		},
		Base: &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.ProductID) (*pb.ProductResponse, error) {
	product, err := h.repo.GetProduct(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   convertToPBTimestamp(product.CreatedAt),
		},
		Base: &pb.BaseResponse{Code: int32(codes.OK)},
	}, nil
}

// Implement other methods (ListProducts, UpdateProduct, DeleteProduct)
