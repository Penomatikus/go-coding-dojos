package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
)

func Test_CreateCharacter(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	playerRepo := repositorytest.ProvidePlayerRepository(&dbStore)
	err := playerRepo.Create(ctx, &model.Player{
		CreatedAt: time.Now(),
		Name:      "Test",
	})
	if err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	handler := NewCharacterHandler(ctx,
		repositorytest.ProvideCharacterRepository(&dbStore),
		playerRepo,
	)

	jsonData, err := json.Marshal(character.CreateRequest{
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
