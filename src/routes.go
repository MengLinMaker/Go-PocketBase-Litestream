package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func totalAdminsRoute(c echo.Context) error {
	totalAdmins, _ := app.Dao().TotalAdmins()
	return c.String(200, fmt.Sprint(totalAdmins))
}

func litestreamMetricsRoute(c echo.Context) error {
	resp, err := http.Get("http://localhost:9090/metrics")
	if err != nil {
		return c.String(503, "Litestream is not running")
	}
	return c.Stream(200, "text/plain", resp.Body)
}
