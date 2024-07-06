package main

import (
	"context"
	"net/http"

	"github.com/Penomatikus/onionarch/internal"
	"github.com/Penomatikus/onionarch/internal/api/rest/middleware"
)

func myHandlerFunc()
func main() {

	ctx := context.Background()
	app := internal.Initialize(ctx)

	mux := http.NewServeMux()
	middlewares := middleware.Compose(middleware.Auth, middleware.Log, middleware.Metrics)

	baseURL := "/api/v1/fatecore"
	mux.Handle("POST "+baseURL+"/session/new", middlewares(app.SessionHandler.StartSession()))
	mux.Handle("POST "+baseURL+"/session/{sessionid}/join", middlewares(app.SessionHandler.JoinSession()))
	mux.Handle("POST "+baseURL+"/session/{sessionid}/leave", middlewares(app.SessionHandler.LeaveSession()))
	mux.Handle("POST "+baseURL+"/character/new", middlewares(app.CharacterHandler.CreateCharacter()))
	mux.Handle("POST "+baseURL+"/character/{id}/update", middlewares(app.CharacterHandler.UpdateCharacter()))
	mux.Handle("POST "+baseURL+"/player/new", middlewares(app.PlayerHandler.CreatePlayer()))
	mux.Handle("POST "+baseURL+"/session/{sessionid}/notification", middlewares(app.NotificationHandler.SendNotification()))
	mux.Handle("GET "+baseURL+"/session/{sessionid}/notification", middlewares(app.NotificationHandler.CollectNotification()))

	server := &http.Server{
		Handler: mux,
	}

	server.ListenAndServe()
}
