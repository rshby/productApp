package repository

import (
	"context"
	"productApp/app/model/entity"
)

type IProductRepository interface {
	Insert(ctx context.Context, input *entity.Product) (*entity.Product, error)
	GetProductsSortByDate(ctx context.Context, order string) ([]entity.Product, error)
	GetProductSortByPrice(ctx context.Context, order string) ([]entity.Product, error)
	GetProductSortByName(ctx context.Context, order string) ([]entity.Product, error)
}
