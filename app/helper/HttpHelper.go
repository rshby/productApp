package helper

import "net/http"

func CodeToSatatus(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "ok"
	case http.StatusNotFound:
		return "not found"
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusUnauthorized:
		return "unauthorized"
	default:
		return "internal server error"
	}
}
