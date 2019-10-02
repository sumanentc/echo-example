package handlers

import (
	"net/http"
	"vo"

	"github.com/labstack/echo"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	//requestID := uuid.NewV4()
	//c.Logger().Infof("RequestID: %s", requestID)

	resp := vo.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
