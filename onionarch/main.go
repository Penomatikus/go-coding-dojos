package main

import (
	"context"
	"net/http"

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

	server.ListenAndServe()
}
