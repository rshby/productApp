package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"productApp/app/config"
	"productApp/app/logging"
	"time"
)

func LoggerMiddleware(cfg config.IConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logfile := logging.NewLoggerFile(cfg)
		startTime := time.Now()

		// capture request body
		var requestMap map[string]any
		json.Unmarshal(ctx.Body(), &requestMap)

		ctx.Next()

		// capture response body
		var responseMap map[string]any
		json.Unmarshal(ctx.Response().Body(), &responseMap)

		// encode request_body and response_body
		requestJson, _ := json.Marshal(&requestMap)
		responseJson, _ := json.Marshal(&responseMap)

		logfile.WithFields(logrus.Fields{
			"url":           string(ctx.Request().URI().Path()),
			"method":        string(ctx.Request().Header.Method()),
			"status_code":   ctx.Response().StatusCode(),
			"response_time": time.Since(startTime).Milliseconds(),
			"request":       string(requestJson),
			"response":      string(responseJson),
		}).Info("incoming request")
		return nil
	}
}
