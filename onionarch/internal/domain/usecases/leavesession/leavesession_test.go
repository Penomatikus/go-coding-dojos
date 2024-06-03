package leavesession

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_Leave_Succes(t *testing.T) {
	db := repositorytest.NewDBStore()

	sessionID := model.SessionID("1337")
	db.Session[sessionID] = &model.Session{
		ID:        "1337",
		Owner:     1337,
		CreatedAt: time.Now(),
		Title:     "Test Session",
	}

	charID := model.CharacterID(1)
	db.Character[charID] = &model.Character{
		ID:        1,
		SessionID: &sessionID,
		PlayerID:  1,
		Points:    100,
	}

	ports := Ports{
		SessionRepo:   repositorytest.ProvideSessionRepository(&db),
		CharacterRepo: repositorytest.ProvideCharacterRepository(&db),
	}

	err := Leave(context.Background(), ports, Request{
		SessionID:   func() model.SessionID { return sessionID },
		CharacterID: func() model.CharacterID { return charID },
	})

	if err != nil {
		t.Fatal(err)
	}

	if db.Character[charID].SessionID != nil {
		t.Fatalf("id was %d expected %s", db.Character[charID].SessionID, "nil")
	}
}
