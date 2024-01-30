package main

import (
	"fmt"
	"runtime/debug"

	"server.bin/framework"

	echo "github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

var app = framework.New()

func main() {
	app.AddRoutes(func(e *core.ServeEvent) {
		e.Router.GET("/hello", adminIdRoute)
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
