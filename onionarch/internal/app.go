package internal

import (
	"context"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest/handler"
	"github.com/Penomatikus/onionarch/internal/infrastructure/db"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"
	"github.com/Penomatikus/onionarch/internal/infrastructure/sessionid"
)

type app struct {
	characterHandler    *handler.CharacterHandler
	notificationHandler *handler.NotificationHandler
	sessionHandler      *handler.SessionHandler
	doActionHandler     *handler.DoActionHandler
}

func Initialize(ctx context.Context) *app {

	// infrastructure
	dbStore := db.NewDBStore()
	notificationService := notification.PrivideService()
	sessionIDGen := sessionid.ProvideSessionIDGen()

	// domain
	characterRepo := db.ProvideCharacterRepository(&dbStore)
	sessionRepo := db.ProvideSessionRepository(&dbStore)

	// handlers
	characterHandler := handler.NewCharacterHandler(ctx, characterRepo)
	doActionHandler := handler.NewDoActionHandler(ctx, characterRepo, sessionRepo)
	notificationHanlder := handler.NewNotificationHandler(ctx, notificationService)
	sessionHandler := handler.NewSessionHandler(ctx, characterRepo, sessionIDGen, sessionRepo)

	return &app{
		characterHandler:    characterHandler,
		notificationHandler: notificationHanlder,
		sessionHandler:      sessionHandler,
		doActionHandler:     doActionHandler,
	}
}

func NewRouterV1(app *app) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /session/new", app.sessionHandler.StartSession)
	router.HandleFunc("POST /session/{sessionid}/join", app.sessionHandler.JoinSession)
	router.HandleFunc("POST /session/{sessionid}/leave", app.sessionHandler.LeaveSession)
	router.HandleFunc("POST /character/new", app.characterHandler.CreateCharacter)
	router.HandleFunc("POST /character/do", app.doActionHandler.DoAction)
	router.HandleFunc("POST /character/{id}/update", app.characterHandler.UpdateCharacter)
	router.HandleFunc("POST /session/{sessionid}/notification", app.notificationHandler.SendNotification)
	router.HandleFunc("GET /session/{sessionid}/notification", app.notificationHandler.CollectNotification)

	base := "/api/v1/fatecore"
	v1 := http.NewServeMux()
	v1.Handle(base+"/", http.StripPrefix(base, router))

	return v1
}
