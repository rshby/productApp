package repository

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"productApp/app/customError"
	"productApp/app/model/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

// function provider
func NewProductRepository(db *sql.DB) IProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

// method insert new data to database
func (p *ProductRepository) Insert(ctx context.Context, input *entity.Product) (*entity.Product, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProductRepository")
	defer span.Finish()

	// prepare queryes
	statement, err := p.DB.PrepareContext(ctxTracing, "INSERT INTO products (name, price, description, quantity) VALUES (?, ?, ?, ?)")
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute query
	result, err := statement.ExecContext(ctxTracing, input.Name, input.Price, input.Description, input.Quantity)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// cek row affected
	if row, _ := result.RowsAffected(); row == 0 {
		return nil, customError.NewInternalSeverError("failed insert data to database")
	}

	// get last insert id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// create response
	input.Id = int(id)

	// return response
	return input, nil
}
