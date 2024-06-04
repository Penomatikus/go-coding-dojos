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
	db.Session["1337"] = &model.Session{
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
		SessionRepo:   repositorytest.ProvideSessionRepository(&db),
		CharacterRepo: repositorytest.ProvideCharacterRepository(&db),
	}

	err := Join(context.Background(), ports, Request{
		SessionID:   func() model.SessionID { return "1337" },
		CharacterID: 1,
	})

	if err != nil {
		t.Fatal(err)
	}

	character, err := ports.CharacterRepo.FindByID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if character.ID != 1 {
		t.Fatalf("id was %d expected %d", character.ID, 1)
	}

	if character.PlayerID != 1 {
		t.Fatalf("id was %d expected %d", character.ID, 1)
	}

	if character.Points != 100 {
		t.Fatalf("id was %d expected %d", character.ID, 100)
	}
}
