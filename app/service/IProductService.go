package service

import (
	"context"
	"productApp/app/model/dto"
)

type IProductService interface {
	AddProduct(ctx context.Context, request *dto.CreateProductRequest) (*dto.CreateProductResponse, error)
	GetProducts(ctx context.Context, sort string, order string) ([]dto.ProductDetailResponse, error)
}
