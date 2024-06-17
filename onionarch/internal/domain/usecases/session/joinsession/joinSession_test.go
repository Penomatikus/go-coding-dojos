package joinsession

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_JoinSession_Success(t *testing.T) {
	db := repositorytest.NewDBStore()

	sessionID := model.SessionID("1337")
	db.Session[sessionID] = &model.Session{
		ID:        "1337",
		Owner:     1337,
		CreatedAt: time.Now(),
		Title:     "Test Session",
	}

	charID := 1
	db.Character[charID] = &model.Character{
		ID:       1,
		PlayerID: 1,
		Points:   100,
	}

	ports := Ports{
		SessionRepository:   repositorytest.ProvideSessionRepository(&db),
		CharacterRepository: repositorytest.ProvideCharacterRepository(&db),
	}

	err := Join(context.Background(), ports, Request{
		SessionID:   model.SessionID("1337"),
		CharacterID: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	character, err := ports.CharacterRepository.FindByID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if *character.SessionID != sessionID {
		t.Fatalf("id was %v expected %v", character.SessionID, sessionID)
	}

}
