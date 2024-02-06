package repository

import (
	"context"
	"productApp/app/model/entity"
)

type IProductRepository interface {
	Insert(ctx context.Context, input *entity.Product) // TODO : lengkapi return nya belum
}
