package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"server.bin/framework"

	echo "github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

var app = framework.New()

func main() {
	app.AddRoutes(func(e *core.ServeEvent) {
		e.Router.GET("/hello", adminIdRoute)
		requireAdmin := apis.RequireAdminAuth()
		e.Router.GET("/litestream", litestreamMetricsRoute, requireAdmin)
	})
	app.Start()
	// Allow Litestream to capture all WAL
	app.DB().NewQuery(`PRAGMA wal_autocheckpoint = 0;`).Execute()
	// Limit Sqlite3 cache to 128MB RAM
	app.DB().NewQuery(`PRAGMA cache_size = -131072000;`).Execute()
	// Limit Go to 100MB RAM
	debug.SetMemoryLimit(100000000)
}

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
