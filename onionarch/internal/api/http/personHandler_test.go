package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	createplayer "github.com/Penomatikus/onionarch/internal/domain/usecases/player/createPlayer"
)

func Test_CreatePerson_Success(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	handler := ProvidePlayerHandler(ctx, repositorytest.ProvidePlayerRepository(&dbStore))
	jsonData, err := json.Marshal(createplayer.Request{Name: "Maggus"})
	if err != nil {
		t.Fatalf("%s: Error marshalling data to JSON", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/fatecore/player/new", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()

	handler.CreatePlayer(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

}
