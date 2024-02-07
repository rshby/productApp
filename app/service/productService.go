package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"productApp/app/customError"
	"productApp/app/helper"
	"productApp/app/model/dto"
	"productApp/app/model/entity"
	"productApp/app/repository"
	"slices"
	"strings"
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

// method list products
func (p *ProductService) GetProducts(ctx context.Context, sort string, order string) ([]dto.ProductDetailResponse, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "Service GetProducts")
	defer span.Finish()

	orders := []string{"ASC", "DESC"}
	if !slices.Contains(orders, strings.ToUpper(order)) {
		return nil, customError.NewBadRequestError("order parameter not in options")
	}

	var products []entity.Product
	// switch by sort
	switch sort {
	case "date":
		// call method in repository
		date, err := p.ProductRepo.GetProductsSortByDate(ctxTracing, strings.ToUpper(order))
		if err != nil {
			return nil, err
		}
		products = date
	case "price":
		// call method in repository
		price, err := p.ProductRepo.GetProductSortByPrice(ctxTracing, strings.ToUpper(order))
		if err != nil {
			return nil, err
		}

		products = price
	case "name":
		// call method in repository
		name, err := p.ProductRepo.GetProductSortByName(ctxTracing, strings.ToUpper(order))
		if err != nil {
			return nil, err
		}

		products = name
	default:
		return nil, customError.NewBadRequestError("sort parameter not in options")
	}

	var response []dto.ProductDetailResponse
	for _, product := range products {
		response = append(response, dto.ProductDetailResponse{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Quantity:    product.Quantity,
			CreatedAt:   helper.DateTimeToString(product.CreatedAt),
			UpdatedAt:   helper.DateTimeToString(product.UpdatedAt),
		})
	}

	// if not found
	if len(response) == 0 {
		return nil, customError.NewNotFoundError("record products not found")
	}

	// success
	return response, nil
}
