package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid/sessionidtest"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

func Test_Session_Success(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dbStrore := repositorytest.NewDBStore()
	sessionIdGen := sessionidtest.ProvideSessionIDGen()

	sessioenRepo := repositorytest.ProvideSessionRepository(&dbStrore)

	characterRepo := repositorytest.ProvideCharacterRepository(&dbStrore)
	if err := characterRepo.Create(ctx, &model.Character{
		Name: "Test",
	}); err != nil {
		t.Fatalf("%s: Error creating character", err)
	}

	notificationPublisher := notification.NewEventBus()
	eventSubscriber := make(notification.EventSubscriber)
	handler := NewSessionHandler(ctx,
		notificationPublisher,
		eventSubscriber,
		characterRepo,
		sessionIdGen,
		sessioenRepo,
	)

	go notification.NewEventSink().Consum(ctx, eventSubscriber)

	var sessionID *string
	t.Run("Start session", func(t *testing.T) {
		jsonData, err := json.Marshal(session.StartRequest{
			Title: "Test",
			Owner: 1,
		})

		if err != nil {
			t.Fatalf("%s: Error marshalling data to JSON", err)
		}

		req := httptest.NewRequest("POST", "/api/v1/fatecore/session/new", bytes.NewReader(jsonData))
		rec := httptest.NewRecorder()

		handler.StartSession(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatalf("expected 200 got %d", res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil || len(data) == 0 {
			t.Fatalf("%s: Error reading response data", err)
		}

		s := string(data)
		sessionID = &s
	})

	t.Run("Join session", func(t *testing.T) {
		jsonData, err := json.Marshal(session.JoinRequest{
			SessionID:   model.SessionID(*sessionID),
			CharacterID: 1,
		})

		if err != nil {
			t.Fatalf("%s: Error marshalling data to JSON", err)
		}

		req := httptest.NewRequest("POST", "/api/v1/fatecore/session/forThisObsolete/join", bytes.NewReader(jsonData))
		req.SetPathValue("sessionid", *sessionID)
		rec := httptest.NewRecorder()

		handler.JoinSession(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatalf("expected 200 got %d", res.StatusCode)
		}

		c, err := characterRepo.FindByID(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if *c.SessionID != model.SessionID(*sessionID) {
			t.Fatal("character did not join")
		}
	})

	t.Run("Leave session", func(t *testing.T) {
		jsonData, err := json.Marshal(session.LeaveRequest{
			SessionID:   model.SessionID(*sessionID),
			CharacterID: 1,
		})

		if err != nil {
			t.Fatalf("%s: Error marshalling data to JSON", err)
		}

		req := httptest.NewRequest("POST", "/api/v1/fatecore/session/forThisObsolete/leave", bytes.NewReader(jsonData))
		req.SetPathValue("sessionid", *sessionID)
		rec := httptest.NewRecorder()

		handler.LeaveSession(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Fatalf("expected 200 got %d", res.StatusCode)
		}

		c, err := characterRepo.FindByID(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if c.SessionID != nil {
			t.Fatal("character did not leave")
		}
	})
}
