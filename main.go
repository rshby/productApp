package main

import (
	"github.com/opentracing/opentracing-go"
	"productApp/app/config"
	"productApp/app/logging"
	"productApp/app/tracing"
	"productApp/database"
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
	database.ConnectDB(cfg, logConsole)

	logConsole.Info(cfg.GetConfig().Logging)
}
