package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func adminIdRoute(c echo.Context) error {
	a, err := app.Dao().FindAdminByEmail("menglinmaker@gmail.com")
	if err != nil {
		fmt.Print("error", err)
	}
	return c.String(200, a.Id)
}

func litestreamMetricsRoute(c echo.Context) error {
	resp, err := http.Get("http://localhost:9090/metrics")
	if err != nil {
		return c.String(503, "Litestream is not running")
	}
	return c.Stream(200, "text/plain", resp.Body)
}
