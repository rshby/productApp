package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"productApp/app/config"
	"productApp/app/logging"
	"productApp/app/tracing"
	"productApp/database"
	"productApp/server"
)

func main() {
	// load config
	cfg := config.NewConfigApp()

	// load log
	logConsole := logging.NewLoggerConsole()

	// set tracing
	tracer, closer := tracing.ConnectJaeger(cfg, logConsole, "productApp")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// connect to DB
	db := database.ConnectDB(cfg, logConsole)

	// create validate
	validate := validator.New()

	// run server
	appServer := server.NewAppServer(cfg, validate, db)
	if err := appServer.RunServer(); err != nil {
		logConsole.Fatalf("cant run server : %v", err)
	}
}
