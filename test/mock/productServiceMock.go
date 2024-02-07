package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"productApp/app/model/dto"
)

type ProductServiceMock struct {
	Mock mock.Mock
}

// function provider
func NewProductServiceMock() *ProductServiceMock {
	return &ProductServiceMock{mock.Mock{}}
}

// method implementasi add product
func (p *ProductServiceMock) AddProduct(ctx context.Context, request *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	args := p.Mock.Called(ctx, request)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*dto.CreateProductResponse), nil
}

// method implementasi get list products
func (p *ProductServiceMock) GetProducts(ctx context.Context, sort string, order string) ([]dto.ProductDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}
