package character

import (
	"context"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_UpdateCharacter_Success(t *testing.T) {
	ctx := context.Background()
	dbStore := repositorytest.NewDBStore()

	characterRepo := repositorytest.ProvideCharacterRepository(&dbStore)
	charID, err := characterRepo.Create(ctx, &model.Character{
		Name:        "Tester",
		Description: "Möp",
	})
	if err != nil {
		t.Fatalf("%s: Error creating chracacter", err)
	}

	ports := UpdatePorts{
		CharacterRepository: characterRepo,
	}
	points := 99
	sId := model.SessionID("1")
	err = Update(context.Background(), ports, UpdateRequest{
		ID:        charID,
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

	if character.Points != points {
		t.Fatalf("points were %d expected %d", character.Points, points)
	}

	if character.SessionID != &sId {
		t.Fatalf("sessionId was %d expected %s", character.SessionID, sId)
	}
}
