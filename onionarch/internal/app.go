package internal

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest/handler"
	"github.com/Penomatikus/onionarch/internal/domain/model"
	domainNotification "github.com/Penomatikus/onionarch/internal/domain/notification"
	"github.com/Penomatikus/onionarch/internal/infrastructure/db"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"

	"github.com/Penomatikus/onionarch/internal/infrastructure/sessionid"
)

type app struct {
	characterHandler     *handler.CharacterHandler
	notificationHandler  *handler.NotificationHandler
	sessionHandler       *handler.SessionHandler
	doHandler            *handler.DoHandler
	notificationConsumer domainNotification.Consumer
	notificationChan     chan model.Notification
}

func Initialize(ctx context.Context) *app {

	// infrastructure
	dbStore := db.NewDBStore()
	sessionIDGen := sessionid.ProvideSessionIDGen()
	notifactionChan := make(chan model.Notification)
	notificationPublisher := notification.NewEventBus()
	notificationConsumer := notification.NewEventSink()

	// domain
	characterRepo := db.ProvideCharacterRepository(&dbStore)
	sessionRepo := db.ProvideSessionRepository(&dbStore)

	// handlers
	characterHandler := handler.NewCharacterHandler(ctx, characterRepo)
	doHandler := handler.NewDoHandler(ctx, characterRepo, sessionRepo, notificationPublisher)
	notificationHandler := handler.NewNotificationHandler(ctx, notificationConsumer)
	sessionHandler := handler.NewSessionHandler(ctx,
		notificationPublisher,
		notifactionChan,
		characterRepo,
		sessionIDGen,
		sessionRepo,
	)

	return &app{
		characterHandler:     characterHandler,
		notificationHandler:  notificationHandler,
		sessionHandler:       sessionHandler,
		doHandler:            doHandler,
		notificationConsumer: notificationConsumer,
		notificationChan:     notifactionChan,
	}
}

func (a *app) Start(ctx context.Context, server *http.Server) {
	log.Println("Starting Server...")
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Println("Starting notification consumer...")
	go func() {
		if err := a.notificationConsumer.Consume(ctx, a.notificationChan); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Println("UP!")
}

func (a *app) Stop(ctx context.Context, server *http.Server) {
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}

	log.Println("Bye.")
}

func NewRouterV1(app *app) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /session/new", app.sessionHandler.StartSession)
	router.HandleFunc("POST /session/{sessionid}/join", app.sessionHandler.JoinSession)
	router.HandleFunc("POST /session/{sessionid}/leave", app.sessionHandler.LeaveSession)
	router.HandleFunc("GET /session/{sessionid}/party", app.sessionHandler.LookupSession)
	router.HandleFunc("GET /session/{sessionid}/notification", app.notificationHandler.CollectNotification)
	router.HandleFunc("POST /character/new", app.characterHandler.CreateCharacter)
	router.HandleFunc("POST /character/do/action", app.doHandler.DoAction)
	router.HandleFunc("POST /character/do/points", app.doHandler.DoPoints)

	base := "/api/v1/fatecore"
	v1 := http.NewServeMux()
	v1.Handle(base+"/", http.StripPrefix(base, router))

	return v1
}
