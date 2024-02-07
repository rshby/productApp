package test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"productApp/app/customError"
	"productApp/app/handler"
	"productApp/app/helper"
	"productApp/app/model/dto"
	mck "productApp/test/mock"
	"strings"
	"testing"
)

func TestAddProductHandler(t *testing.T) {
	t.Run("test add product handler error validasi", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Post("/", productHandler.AddProduct)

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test add product error internal server", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Post("/", productHandler.AddProduct)

		// mock service
		productService.Mock.On("AddProduct", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError("error cant failed"))

		// test
		requestBody := dto.CreateProductRequest{
			Name:        "iPhone 15",
			Price:       22000000,
			Description: "hp",
			Quantity:    100,
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create http Request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive respone
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, helper.CodeToSatatus(http.StatusInternalServerError), responseBody["status"].(string))
	})
	t.Run("test add product error not found", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Post("/", productHandler.AddProduct)

		// mock
		errorMessage := "record not found"
		productService.Mock.On("AddProduct", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		requestBody := dto.CreateProductRequest{
			Name:        "iPhone 12",
			Price:       12000000,
			Description: "hp",
			Quantity:    100,
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusNotFound, response.StatusCode)

		// receive response body
		body, err := io.ReadAll(response.Body)
		assert.Nil(t, err)

		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
	t.Run("test add product error bad request", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Post("/", productHandler.AddProduct)

		// mock
		errorMessage := "error bad request"
		productService.Mock.On("AddProduct", mock.Anything, mock.Anything).
			Return(nil, customError.NewBadRequestError(errorMessage), validator.ValidationErrors{})

		// test
		requestBody := dto.CreateProductRequest{
			Name:        "iPhone 12",
			Price:       12000000,
			Description: "hp",
			Quantity:    10,
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)

		// get response body
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
	t.Run("test add product success insert", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Post("/", productHandler.AddProduct)

		// mock
		productService.Mock.On("AddProduct", mock.Anything, mock.Anything).
			Return(&dto.CreateProductResponse{
				Id:          1,
				Name:        "iPhone 12",
				Price:       12000000,
				Description: "hp",
				Quantity:    10,
				CreatedAt:   "2020-10-10 10:00:00",
				UpdatedAt:   "2020-10-10 10:00:00",
			}, nil)

		// test
		requestBody := dto.CreateProductRequest{
			Name:        "iPhone 12",
			Price:       12000000,
			Description: "hp",
			Quantity:    10,
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		// receive response body
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, "ok", responseBody["status"].(string))
	})
}

func TestGetProductsHandler(t *testing.T) {
	t.Run("test get products error not found", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		// mock
		errorMessage := "record products not found"
		productService.Mock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		// create request
		request := httptest.NewRequest(http.MethodGet, "/products?sort=name&order=asc", nil)

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusNotFound, response.StatusCode)

		// get responsebody
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, "not found", responseBody["status"].(string))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
	t.Run("test get products error bad request", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		// mock
		errorMessage := "error bad request"
		productService.Mock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, customError.NewBadRequestError(errorMessage))

		// test
		request := httptest.NewRequest(http.MethodGet, "/products", nil)

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)

		// get responsebody
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, "bad request", responseBody["status"].(string))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
	t.Run("test get products error internal server error", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		// mock
		errorMessage := "error internal server"
		productService.Mock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		request := httptest.NewRequest(http.MethodGet, "/products", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

		// get responsebody
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
	t.Run("test get products success", func(t *testing.T) {
		productService := mck.NewProductServiceMock()
		productHandler := handler.NewProductHandler(productService)

		app := fiber.New()
		app.Get("/products", productHandler.GetProducts)

		// mock
		productService.Mock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).
			Return([]dto.ProductDetailResponse{
				{
					Id:          1,
					Name:        "iPhone 12",
					Price:       12000000,
					Description: "hp",
					Quantity:    10,
					CreatedAt:   "2020-10-10 10:00:00",
					UpdatedAt:   "2020-10-10 10:00:00",
				},
			}, nil)

		// test
		request := httptest.NewRequest(http.MethodGet, "/products?sort=price&order=asc", nil)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		// get responsebody
		body, err := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Nil(t, err)
		assert.Equal(t, "ok", responseBody["status"].(string))
	})
}
