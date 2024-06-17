package updatecharacter

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_UpdateCharacter_Success(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	playerRepo := repositorytest.ProvidePlayerRepository(&dbStore)
	if err := playerRepo.Create(ctx, &model.Player{
		CreatedAt: time.Now(),
		Name:      "Test",
	}); err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	characterRepo := repositorytest.ProvideCharacterRepository(&dbStore)
	if err := characterRepo.Create(ctx, &model.Character{
		Name:        "Tester",
		Description: "Möp",
		PlayerID:    1,
	}); err != nil {
		t.Fatalf("%s: Error creating player", err)
	}

	ports := Ports{
		PlayerRepository:    playerRepo,
		CharacterRepository: characterRepo,
	}
	points := 99
	sId := model.SessionID("1")
	err := Update(context.Background(), ports, Request{
		ID:        1,
		PlayerID:  1,
		Points:    &points,
		SessionID: &sId,
	})

	if err != nil {
		t.Fatal(err)
	}

	character, err := ports.CharacterRepository.FindByID(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if character.ID != 1 {
		t.Fatalf("id was %d expected %d", character.ID, 1)
	}

	if character.PlayerID != 1 {
		t.Fatalf("playerId was %d expected %d", character.PlayerID, 1)
	}

	if character.Points != points {
		t.Fatalf("points were %d expected %d", character.Points, points)
	}

	if character.SessionID != &sId {
		t.Fatalf("sessionId was %d expected %s", character.SessionID, sId)
	}
}
