package main

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v5"
)

func timerShutdown(timer *time.Timer, timeout time.Duration) {
	timer.Stop()
	timer.Reset(timeout)
	<-timer.C
	fmt.Println("Shutdown server - idle for", timeout)
	defer os.Exit(0)
}

func idleShutdown(timeout time.Duration) echo.MiddlewareFunc {
	timer := time.NewTimer(timeout)
	go timerShutdown(timer, timeout)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		go timerShutdown(timer, timeout)
		return next
	}
}
