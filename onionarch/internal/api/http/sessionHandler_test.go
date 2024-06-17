package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid/sessionidtest"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session/joinsession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session/leavesession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session/startsession"
)

func Test_Session_Success(t *testing.T) {
	ctx := context.Background()
	dbStrore := repositorytest.NewDBStore()
	sessionIdGen := sessionidtest.ProvideSessionIDGen()

	sessioenRepo := repositorytest.ProvideSessionRepository(&dbStrore)

	playerRepo := repositorytest.ProvidePlayerRepository(&dbStrore)
	if err := playerRepo.Create(ctx, &model.Player{
		CreatedAt: time.Now(),
		Name:      "Test",
	}); err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	characterRepo := repositorytest.ProvideCharacterRepository(&dbStrore)
	if err := characterRepo.Create(ctx, &model.Character{
		Name:     "Test",
		PlayerID: 1,
	}); err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	handler := ProvidesessionHandler(ctx, characterRepo, playerRepo, sessionIdGen, sessioenRepo)

	var sessionID *string
	t.Run("Start session", func(t *testing.T) {
		jsonData, err := json.Marshal(startsession.Request{
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
		jsonData, err := json.Marshal(joinsession.Request{
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
		jsonData, err := json.Marshal(leavesession.Request{
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
