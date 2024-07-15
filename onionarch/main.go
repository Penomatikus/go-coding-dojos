package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Penomatikus/onionarch/internal"
	"github.com/Penomatikus/onionarch/internal/api/rest/middleware"
)

func main() {

	ctx := context.Background()
	app := internal.Initialize(ctx)
	router := internal.NewRouterV1(app)
	middlewares := middleware.Compose(middleware.Auth, middleware.Log, middleware.Metrics)

	server := &http.Server{
		Handler: middlewares(router),
	}

	log.Println("Starting...")
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(ctx, 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	log.Println("Bye.")
}
