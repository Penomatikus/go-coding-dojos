package createcharacter

import (
	"context"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_CreateCharacter_Success(t *testing.T) {
	db := repositorytest.NewDBStore()
	db.Player[1] = &model.Player{
		ID:        1,
		CreatedAt: time.Now(),
		Name:      "Ingeborg",
	}

	ports := Ports{
		playerRepository:    repositorytest.ProvidePlayerRepository(&db),
		characterRepository: repositorytest.ProvideCharacterRepository(&db),
	}

	err := Create(context.Background(), ports, Request{
		PlayerID:    1,
		Name:        "Wilde Inge",
		Description: "Wild wie zwei Juttas",
	})

	if err != nil {
		t.Fatal(err)
	}

	character, err := ports.characterRepository.FindByID(context.Background(), 1)
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
