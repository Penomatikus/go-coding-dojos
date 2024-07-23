package character

import (
	"context"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_CreateCharacter_Success(t *testing.T) {
	db := repositorytest.NewDBStore()

	ports := CreatePorts{
		CharacterRepository: repositorytest.ProvideCharacterRepository(&db),
	}

	charID, err := Create(context.Background(), ports, CreateRequest{
		Name:        "Wilde Inge",
		Description: "Wild wie zwei Juttas",
	})

	if err != nil {
		t.Fatal(err)
	}

	character, err := ports.CharacterRepository.FindByID(context.Background(), charID)
	if err != nil {
		t.Fatal(err)
	}

	if character.ID != 1 {
		t.Fatalf("id was %d expected %d", character.ID, 1)
	}

	if character.Points != 0 {
		t.Fatalf("points were %d expected %d", character.Points, 0)
	}
}
