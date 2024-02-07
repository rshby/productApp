package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"productApp/app/customError"
	"productApp/app/helper"
	"productApp/app/model/dto"
	"productApp/app/service"
	"strings"
)

type ProductHandler struct {
	ProductService service.IProductService
}

// function provider
func NewProductHandler(productService service.IProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

// method add product
func (p *ProductHandler) AddProduct(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "Handler AddProduct")
	defer span.Finish()

	// parsing body request
	var request dto.CreateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToSatatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call service
	product, err := p.ProductService.AddProduct(ctxTracing, &request)
	if err != nil {
		// if error validation
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessage []string
			for _, fieldError := range validationErrors {
				msg := fmt.Sprintf("error on field [%v], with tag [%v]", fieldError.Field(), fieldError.Tag())
				errorMessage = append(errorMessage, msg)
			}

			statusCode := http.StatusBadRequest
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToSatatus(statusCode),
				Message:    strings.Join(errorMessage, ". "),
			})
		}

		var statusCode int
		switch err.(type) {
		case *customError.NotFoundError:
			statusCode = http.StatusNotFound
		case *customError.BadRequestError:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToSatatus(statusCode),
			Message:    err.Error(),
		})
	}

	// success insert
	statusCode := http.StatusOK
	ctx.Status(statusCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.CodeToSatatus(statusCode),
		Message:    "success insert",
		Data:       product,
	})
}

// method get list products
func (p *ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "Handler GetProducts")
	defer span.Finish()

	// get query
	sort := ctx.Query("sort", "date")
	order := ctx.Query("order", "DESC")

	// call procedure in service
	products, err := p.ProductService.GetProducts(ctxTracing, sort, order)
	if err != nil {
		var statusCode int
		switch err.(type) {
		case *customError.NotFoundError:
			statusCode = http.StatusNotFound
		case *customError.BadRequestError:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToSatatus(statusCode),
			Message:    err.Error(),
		})
	}

	// success get data
	statusCode := http.StatusOK
	ctx.Status(statusCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.CodeToSatatus(statusCode),
		Message:    "success get products",
		Data:       products,
	})
}
