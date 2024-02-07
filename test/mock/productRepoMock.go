package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"productApp/app/model/entity"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

// function provider
func NewProductRepoMock() *ProductRepoMock {
	return &ProductRepoMock{mock.Mock{}}
}

// method mock insert data product to database
func (p *ProductRepoMock) Insert(ctx context.Context, input *entity.Product) (*entity.Product, error) {
	args := p.Mock.Called(ctx, input)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*entity.Product), nil
}

func (p *ProductRepoMock) GetProductsSortByDate(ctx context.Context, order string) ([]entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepoMock) GetProductSortByPrice(ctx context.Context, order string) ([]entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductRepoMock) GetProductSortByName(ctx context.Context, order string) ([]entity.Product, error) {
	//TODO implement me
	panic("implement me")
}
