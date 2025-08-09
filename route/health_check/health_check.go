package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheck(app *echo.Echo) {
	app.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HealthCheckResponse{Message: "OK"})
	})
}
