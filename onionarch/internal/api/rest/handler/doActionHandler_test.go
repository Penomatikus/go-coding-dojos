package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

func Test_DoAction(t *testing.T) {
	ctx := context.Background()
	dbStrore := repositorytest.NewDBStore()
	characterRepo := repositorytest.ProvideCharacterRepository(&dbStrore)
	sessionRepo := repositorytest.ProvideSessionRepository(&dbStrore)

	sessionID := model.SessionID("test")
	err := sessionRepo.Create(ctx, &model.Session{
		ID:        sessionID,
		CreatedAt: time.Now(),
		Title:     "Test",
	})
	if err != nil {
		t.Fatal(err)
	}

	err = characterRepo.Create(ctx, &model.Character{
		Name:      "Tester",
		SessionID: &sessionID,
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := json.Marshal(character.DoRequest{
		ActionName:  "Teste",
		CharacterID: 1,
		Costs:       10,
		SessionID:   sessionID,
	})
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/api/v1/fatecore/character/do", bytes.NewReader(request))
	rec := httptest.NewRecorder()

	notificationPublisher := notification.NewEventBus()
	NewDoActionHandler(ctx, characterRepo, sessionRepo, notificationPublisher).DoAction(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response DoActionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(response)

}
