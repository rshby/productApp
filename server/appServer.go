package server

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"productApp/app/config"
	"productApp/app/handler"
	"productApp/app/middleware"
	"productApp/app/repository"
	"productApp/app/service"
	"productApp/routes"
)

type AppServer struct {
	Router *fiber.App
	Config config.IConfig
}

// function provider
func NewAppServer(cfg config.IConfig, validate *validator.Validate, db *sql.DB) IServer {
	// register repository
	productRepo := repository.NewProductRepository(db)

	// register service
	productService := service.NewProductService(validate, productRepo)

	// register handler
	productHandler := handler.NewProductHandler(productService)

	// middleware
	loggerMiddleware := middleware.LoggerMiddleware(cfg)

	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(logger.New())

	v1 := app.Group("/api/v1").Use(loggerMiddleware)

	// generate routes
	routes.SetProductRoutes(v1, productHandler)

	return &AppServer{
		Router: app,
		Config: cfg,
	}
}

// method implementasi run server
func (a *AppServer) RunServer() error {
	addr := fmt.Sprintf(":%v", a.Config.GetConfig().App.Port)
	return a.Router.Listen(addr)
}
