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
	"github.com/Penomatikus/onionarch/internal/domain/model"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	eventSubscriber := make(chan model.Notification)
	app := internal.Initialize(ctx, eventSubscriber)

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

	go internal.NewNotificationConsumer().Consum(ctx, eventSubscriber)

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
