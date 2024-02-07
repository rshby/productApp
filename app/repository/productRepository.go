package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (p *ProductRepository) GetProductsSortByDate(ctx context.Context, order string) ([]entity.Product, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "Repository GetProductsSortByDate")
	defer span.Finish()

	// create statement
	query := fmt.Sprintf("SELECT id, name, price, description, quantity, created_at, updated_at FROM products ORDER BY created_at %v", order)
	statement, err := p.DB.PrepareContext(ctxTracing, query)
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute
	rows, err := statement.QueryContext(ctxTracing)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	var products []entity.Product
	for rows.Next() {
		// scan data
		var product entity.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.CreatedAt, &product.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, customError.NewNotFoundError("record not found")
			}

			return nil, customError.NewInternalSeverError(err.Error())
		}

		// append
		products = append(products, product)
	}

	// cek jika data not found
	if len(products) == 0 {
		return nil, customError.NewNotFoundError("record products not found")
	}

	// success
	return products, nil
}

func (p *ProductRepository) GetProductSortByPrice(ctx context.Context, order string) ([]entity.Product, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "Repository GetProductSortByPrice")
	defer span.Finish()

	// create statement
	query := fmt.Sprintf("SELECT id, name, price, description, quantity, created_at, updated_at FROM products ORDER BY price %v", order)
	statement, err := p.DB.PrepareContext(ctxTracing, query)
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute
	rows, err := statement.QueryContext(ctxTracing)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	var products []entity.Product
	for rows.Next() {
		// scan data
		var product entity.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.CreatedAt, &product.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, customError.NewNotFoundError("record not found")
			}

			return nil, customError.NewInternalSeverError(err.Error())
		}

		// append
		products = append(products, product)
	}

	// cek jika data not found
	if len(products) == 0 {
		return nil, customError.NewNotFoundError("record products not found")
	}

	// success
	return products, nil
}

func (p *ProductRepository) GetProductSortByName(ctx context.Context, order string) ([]entity.Product, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "Repository GetProductSortByName")
	defer span.Finish()

	// create statement
	query := fmt.Sprintf("SELECT id, name, price, description, quantity, created_at, updated_at FROM products ORDER BY name %v", order)
	statement, err := p.DB.PrepareContext(ctxTracing, query)
	defer statement.Close()
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	// execute
	rows, err := statement.QueryContext(ctxTracing)
	if err != nil {
		return nil, customError.NewInternalSeverError(err.Error())
	}

	var products []entity.Product
	for rows.Next() {
		// scan data
		var product entity.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.Quantity, &product.CreatedAt, &product.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, customError.NewNotFoundError("record not found")
			}

			return nil, customError.NewInternalSeverError(err.Error())
		}

		// append
		products = append(products, product)
	}

	// cek jika data not found
	if len(products) == 0 {
		return nil, customError.NewNotFoundError("record products not found")
	}

	// success
	return products, nil
}
