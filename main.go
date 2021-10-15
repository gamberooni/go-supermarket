package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gamberooni/go-supermarket/router"
	"github.com/gamberooni/go-supermarket/util"
)

func main() {
	// initialize db
	db := util.InitDB()

	// create router using the sqlite db
	r := router.New(db)

	// Start server
	go func() {
		if err := r.Start(":1323"); err != nil && err != http.ErrServerClosed {
			r.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}
}
