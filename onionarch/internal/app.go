package internal

import (
	"context"

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

// https://www.alexedwards.net/blog/making-and-using-middleware
// https://surajincloud.com/understanding-http-server-in-go-mux
// func NewMux() *http.ServeMux {

// }
