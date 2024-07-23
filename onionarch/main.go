package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Penomatikus/onionarch/internal"
	"github.com/Penomatikus/onionarch/internal/api/rest/middleware"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "Port number to listen on (default 8080)")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := internal.Initialize(ctx)
	router := internal.NewRouterV1(app)

	flag.Parse()
	if port == 0 {
		port = 8080
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware.Compose(middleware.Auth, middleware.Log, middleware.Metrics)(router),
	}

	app.Start(ctx, server)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("\nReceived shudown signal")

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	app.Stop(ctx, server)
}
