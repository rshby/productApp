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
func NewProductServiceMock() *ProductRepoMock {
	return &ProductRepoMock{mock.Mock{}}
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
