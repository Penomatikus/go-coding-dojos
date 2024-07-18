package internal

import (
	"context"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest/handler"
	domainNotification "github.com/Penomatikus/onionarch/internal/domain/notification"
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

func Initialize(ctx context.Context, eventSubscriber notification.EventSubscriber) *app {

	// infrastructure
	dbStore := db.NewDBStore()
	sessionIDGen := sessionid.ProvideSessionIDGen()
	notificationPublisher := notification.NewEventBus()
	notificationConsumer := notification.NewEventSink()

	// domain
	characterRepo := db.ProvideCharacterRepository(&dbStore)
	sessionRepo := db.ProvideSessionRepository(&dbStore)

	// handlers
	characterHandler := handler.NewCharacterHandler(ctx, characterRepo)
	doActionHandler := handler.NewDoActionHandler(ctx, characterRepo, sessionRepo, notificationPublisher)
	notificationHandler := handler.NewNotificationHandler(ctx, notificationConsumer)
	sessionHandler := handler.NewSessionHandler(ctx,
		notificationPublisher,
		eventSubscriber,
		characterRepo,
		sessionIDGen,
		sessionRepo,
	)

	return &app{
		characterHandler:    characterHandler,
		notificationHandler: notificationHandler,
		sessionHandler:      sessionHandler,
		doActionHandler:     doActionHandler,
	}
}

func NewNotificationConsumer() domainNotification.Consumer {
	return notification.NewEventSink()
}

func NewRouterV1(app *app) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /session/new", app.sessionHandler.StartSession)
	router.HandleFunc("POST /session/{sessionid}/join", app.sessionHandler.JoinSession)
	router.HandleFunc("POST /session/{sessionid}/leave", app.sessionHandler.LeaveSession)
	router.HandleFunc("POST /character/new", app.characterHandler.CreateCharacter)
	router.HandleFunc("POST /character/do", app.doActionHandler.DoAction)
	router.HandleFunc("POST /character/{id}/update", app.characterHandler.UpdateCharacter)
	router.HandleFunc("GET /session/{sessionid}/notification", app.notificationHandler.CollectNotification)

	base := "/api/v1/fatecore"
	v1 := http.NewServeMux()
	v1.Handle(base+"/", http.StripPrefix(base, router))

	return v1
}
