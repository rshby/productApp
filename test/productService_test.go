package test

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"productApp/app/customError"
	"productApp/app/model/dto"
	"productApp/app/model/entity"
	"productApp/app/service"
	mck "productApp/test/mock"
	"testing"
	"time"
)

func TestAddProduct(t *testing.T) {
	t.Run("test add product error validasi", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// request
		request := dto.CreateProductRequest{
			Name:        "",
			Price:       0,
			Description: "",
			Quantity:    0,
		}

		// call method add product in service
		product, err := productService.AddProduct(context.Background(), &request)
		_, ok := err.(validator.ValidationErrors)

		assert.Nil(t, product)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.True(t, ok)
	})
	t.Run("test add product error failed to insert", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		errorMessage := "failed to insert new data"
		productRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		product, err := productService.AddProduct(context.Background(), &dto.CreateProductRequest{
			Name:        "iPhone 15 Pro Max",
			Price:       24999000,
			Description: "iPhone generasi terbaru",
			Quantity:    100,
		})

		assert.Nil(t, product)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
		productRepo.Mock.AssertExpectations(t)
	})
	t.Run("test add product success", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		productRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Product{
				Id:          1,
				Name:        "Macbook Pro M3 Pro 18/1TB",
				Price:       34900000,
				Description: "Macbook paling baru",
				Quantity:    100,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}, nil)

		// test
		product, err := productService.AddProduct(context.Background(), &dto.CreateProductRequest{
			Name:        "Macbook Pro M3 Pro 18/1TB",
			Price:       34900000,
			Description: "Macbook paling baru",
			Quantity:    100,
		})

		assert.Nil(t, err)
		assert.NotNil(t, product)
	})
}

func TestGetListProducts(t *testing.T) {
	t.Run("get products error order not in options", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// test
		products, err := productService.GetProducts(context.Background(), "date", "QWERTY")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "order parameter not in options", err.Error())
	})
	t.Run("get products error sort not in options", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// test
		products, err := productService.GetProducts(context.Background(), "age", "desc")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "sort parameter not in options", err.Error())
	})
	t.Run("get products sort date error not found", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		errorMessage := "record not found"
		productRepo.Mock.On("GetProductsSortByDate", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		products, err := productService.GetProducts(context.Background(), "date", "asc")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
	t.Run("get products sort date success", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		productRepo.Mock.On("GetProductsSortByDate", mock.Anything, mock.Anything).
			Return([]entity.Product{
				{
					Id:          1,
					Name:        "iPhone 12",
					Price:       12000000,
					Description: "hp",
					Quantity:    100,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}, nil)

		// test
		products, err := productService.GetProducts(context.Background(), "date", "asc")
		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Equal(t, 1, products[0].Id)
	})
	t.Run("get products sort price error not found", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		errorMessage := "record not found"
		productRepo.Mock.On("GetProductSortByPrice", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		products, err := productService.GetProducts(context.Background(), "price", "asc")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("get products sort price success", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		productRepo.Mock.On("GetProductSortByPrice", mock.Anything, mock.Anything).
			Return([]entity.Product{
				{
					Id:          1,
					Name:        "iPhone 12",
					Price:       12000000,
					Description: "hp",
					Quantity:    100,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
				{
					Id:          2,
					Name:        "iPhone 12 Pro",
					Price:       13000000,
					Description: "hp",
					Quantity:    100,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}, nil)

		// test
		products, err := productService.GetProducts(context.Background(), "price", "asc")
		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Equal(t, 2, len(products))
		assert.Equal(t, 1, products[0].Id)
		productRepo.Mock.AssertExpectations(t)
	})
	t.Run("get products sort name error not found", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		errorMessage := "record no found"
		productRepo.Mock.On("GetProductSortByName", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		products, err := productService.GetProducts(context.Background(), "name", "asc")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("get products sort name success", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		productRepo.Mock.On("GetProductSortByName", mock.Anything, mock.Anything).
			Return([]entity.Product{
				{
					Id:          1,
					Name:        "iPhone 12",
					Price:       12000000,
					Description: "hp",
					Quantity:    100,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			}, nil)

		// test
		products, err := productService.GetProducts(context.Background(), "name", "asc")
		assert.Nil(t, err)
		assert.NotNil(t, products)
		assert.Equal(t, 1, products[0].Id)
	})
	t.Run("get products error not found", func(t *testing.T) {
		validate := validator.New()
		productRepo := mck.NewProductRepoMock()
		productService := service.NewProductService(validate, productRepo)

		// mock
		productRepo.Mock.On("GetProductSortByName", mock.Anything, mock.Anything).
			Return([]entity.Product{}, nil)

		// test
		products, err := productService.GetProducts(context.Background(), "name", "asc")
		assert.Nil(t, products)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "record products not found", err.Error())
	})
}
