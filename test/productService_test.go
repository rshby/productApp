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
