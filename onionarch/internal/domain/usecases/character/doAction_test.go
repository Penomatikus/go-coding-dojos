package character

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid/sessionidtest"
)

func Test_DoAction(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	sessionID, err := sessionidtest.ProvideSessionIDGen().GenerateSessionID()
	if err != nil {
		t.Fatal(err)
	}

	sessionRepo := repositorytest.ProvideSessionRepository(&dbStore)
	if err := sessionRepo.Create(ctx, &model.Session{
		Owner:     1,
		CreatedAt: time.Now(),
		ID:        sessionID,
		Title:     "Test",
	}); err != nil {
		t.Fatal(err)
	}

	characterRepo := repositorytest.ProvideCharacterRepository(&dbStore)
	if err := characterRepo.Create(ctx, &model.Character{
		Name:        "Tester",
		Description: "Möp",
		PlayerID:    1,
		SessionID:   &sessionID,
	}); err != nil {
		t.Fatal(err)
	}

	doPorts := DoPorts{
		CharacterRepository: characterRepo,
		SessionRepository:   sessionRepo,
	}

	charPoints, err := Do(ctx, doPorts, DoRequest{ActionName: "Reject PR", CharacterID: 1, Costs: -10, SessionID: sessionID})
	if err != nil {
		t.Fatal(err)
	}

	if charPoints != -10 {
		t.Fatalf("Got %d; Want %d", charPoints, -10)
	}

	charPoints, err = Do(ctx, doPorts, DoRequest{ActionName: "Approve PR", CharacterID: 1, Costs: 10, SessionID: sessionID})
	if err != nil {
		t.Fatal(err)
	}

	if charPoints != 0 {
		t.Fatalf("Got %d; Want %d", charPoints, 0)
	}
}
