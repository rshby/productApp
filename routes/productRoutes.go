package routes

import (
	"github.com/gofiber/fiber/v2"
	"productApp/app/handler"
)

func SetProductRoutes(r fiber.Router, handler *handler.ProductHandler) {
	r.Post("/product", handler.AddProduct)
	r.Get("/products", handler.GetProducts)
}
