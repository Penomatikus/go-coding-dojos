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

	mux := http.NewServeMux()

	middlewares := middleware.Compose(middleware.Auth, middleware.Log, middleware.Metrics)

	mux.Handle("POST /api/v1/fatecore/session/new", middlewares(app.SessionHandler.StartSession()))
	mux.Handle("POST /api/v1/fatecore/session/{sessionid}/join", middlewares(app.SessionHandler.JoinSession()))
	mux.Handle("POST /api/v1/fatecore/session/{sessionid}/leave", middlewares(app.SessionHandler.LeaveSession()))

	mux.Handle("POST /api/v1/fatecore/character/new", middlewares(app.CharacterHandler.CreateCharacter()))
	mux.Handle("POST /api/v1/fatecore/character/{id}/update", middlewares(app.CharacterHandler.UpdateCharacter()))

	mux.Handle("POST /api/v1/fatecore/player/new", middlewares(app.PlayerHandler.CreatePlayer()))

	mux.Handle("POST /api/v1/fatecore/session/{sessionid}/notification", middlewares(app.NotificationHandler.SendNotification()))
	mux.Handle("GET /api/v1/fatecore/session/{sessionid}/notification", middlewares(app.NotificationHandler.CollectNotification()))

	server := &http.Server{
		Handler: mux,
	}

	server.ListenAndServe()
}
