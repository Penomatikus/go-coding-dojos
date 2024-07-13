package internal

import (
	"context"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/infrastructure/db"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"
	"github.com/Penomatikus/onionarch/internal/infrastructure/sessionid"
)

type app struct {
	CharacterHandler    *rest.CharacterHandler
	NotificationHandler *rest.NotificationHandler
	PlayerHandler       *rest.PlayerHandler
	SessionHandler      *rest.SessionHandler
}

func Initialize(ctx context.Context) *app {

	// infrastructure
	dbStore := db.NewDBStore()
	notificationService := notification.PrivideService()
	sessionIDGen := sessionid.ProvideSessionIDGen()

	// domain
	characterRepo := db.ProvideCharacterRepository(&dbStore)
	playerRepo := db.ProvidePlayerRepository(&dbStore)
	sessionRepo := db.ProvideSessionRepository(&dbStore)

	// handlers
	characterHandler := rest.NewCharacterHandler(ctx, characterRepo, playerRepo)
	notificationHanlder := rest.NewNotificationHandler(ctx, notificationService)
	playerHandler := rest.NewPlayerHandler(ctx, playerRepo)
	sessionHandler := rest.NewSessionHandler(ctx, characterRepo, playerRepo, sessionIDGen, sessionRepo)

	return &app{
		CharacterHandler:    characterHandler,
		NotificationHandler: notificationHanlder,
		PlayerHandler:       playerHandler,
		SessionHandler:      sessionHandler,
	}
}

func NewRouterV1(app *app) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /session/new", app.SessionHandler.StartSession)
	router.HandleFunc("POST /session/{sessionid}/join", app.SessionHandler.JoinSession)
	router.HandleFunc("POST /session/{sessionid}/leave", app.SessionHandler.LeaveSession)
	router.HandleFunc("POST /character/new", app.CharacterHandler.CreateCharacter)
	router.HandleFunc("POST /character/{id}/update", app.CharacterHandler.UpdateCharacter)
	router.HandleFunc("POST /player/new", app.PlayerHandler.CreatePlayer)
	router.HandleFunc("POST /session/{sessionid}/notification", app.NotificationHandler.SendNotification)
	router.HandleFunc("GET /session/{sessionid}/notification", app.NotificationHandler.CollectNotification)

	base := "/api/v1/fatecore"
	v1 := http.NewServeMux()
	v1.Handle(base+"/", http.StripPrefix(base, router))

	return v1
}
