package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"productApp/app/helper"
	"productApp/app/model/dto"
	"productApp/app/model/entity"
	"productApp/app/repository"
	"time"
)

type ProductService struct {
	Validate    *validator.Validate
	ProductRepo repository.IProductRepository
}

// function provider
func NewProductService(validate *validator.Validate, productRepo repository.IProductRepository) IProductService {
	return &ProductService{
		Validate:    validate,
		ProductRepo: productRepo,
	}
}

// method implement AddProduct to database
func (p *ProductService) AddProduct(ctx context.Context, request *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "Service AddProduct")
	defer span.Finish()

	// validate request
	if err := p.Validate.StructCtx(ctxTracing, *request); err != nil {
		return nil, err
	}

	// create entity
	input := entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Quantity:    request.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// call procedure in repository
	result, err := p.ProductRepo.Insert(ctxTracing, &input)
	if err != nil {
		return nil, err
	}

	// mapping to response dto
	response := dto.CreateProductResponse{
		Id:          result.Id,
		Name:        result.Name,
		Price:       result.Price,
		Description: result.Description,
		Quantity:    result.Quantity,
		CreatedAt:   helper.DateTimeToString(result.CreatedAt),
		UpdatedAt:   helper.DateTimeToString(result.UpdatedAt),
	}

	// return response
	return &response, nil
}
