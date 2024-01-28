package main

import (
	"log"
	"os"
	"runtime/debug"

	echo "github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var app = pocketbase.NewWithConfig(pocketbase.Config{
	DefaultDataDir: "../db",
})

func main() {
	debug.SetGCPercent(10000)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello", hello)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	//app.Dao().DB().NewQuery(`PRAGMA journal_mode = WAL;`).SQL()
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Cannot read file: " + path)
	}
	return string(data)
}

func hello(c echo.Context) error {
	return c.String(200, "Hello world!")
}
