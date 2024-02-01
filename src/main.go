package main

import (
	"runtime/debug"
	"time"

	"server.bin/framework"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

var app = framework.New()

func main() {
	app.AddRoutes(func(e *core.ServeEvent) {
		e.Router.Use(
			apis.ActivityLogger(app.Pb),
			// Scale idle server to zero
			// Covers hourly Litestream backup
			idleShutdown(75*time.Minute),
		)
		e.Router.GET("/total_admins", totalAdminsRoute)
		e.Router.GET("/litestream", litestreamMetricsRoute, apis.RequireAdminAuth())
	}).Start()
	// Allow Litestream to capture all WAL
	app.DB().NewQuery(`PRAGMA wal_autocheckpoint = 0;`).Execute()
	// Limit Sqlite3 cache to 128MB RAM
	app.DB().NewQuery(`PRAGMA cache_size = -131072000;`).Execute()
	// Limit Go to 100MB RAM
	debug.SetMemoryLimit(100000000)
}
