package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	createcharacter "github.com/Penomatikus/onionarch/internal/domain/usecases/createCharacter"
	"github.com/Penomatikus/onionarch/internal/infrastructure/db"
)

func Test_CreateCharacter(t *testing.T) {
	ctx := context.Background()
	dbStore := db.NewDBStore()

	playerRepo := db.ProvidePlayerRepository(&dbStore)
	err := playerRepo.Create(ctx, &model.Player{
		CreatedAt: time.Now(),
		Name:      "Test",
	})
	if err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	handler := characterHandler{
		ctx:                 ctx,
		characterRepository: db.ProvideCharacterRepository(&dbStore),
		playerRepository:    playerRepo,
	}

	jsonData, err := json.Marshal(createcharacter.Request{
		PlayerID:    1,
		Name:        "Hallo",
		Description: "Test",
	})

	if err != nil {
		t.Fatalf("%s: Error marshalling data to JSON", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/fatecore/character/new", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()

	handler.CreateCharacter(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

}
