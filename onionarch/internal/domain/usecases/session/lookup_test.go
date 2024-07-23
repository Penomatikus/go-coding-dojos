package session

import (
	"context"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository/repositorytest"
)

func Test_Lookup(t *testing.T) {
	db := repositorytest.NewDBStore()

	ports := LookupPorts{
		SessionRepository:   repositorytest.ProvideSessionRepository(&db),
		CharacterRepository: repositorytest.ProvideCharacterRepository(&db),
	}

	sessionID := model.SessionID("test")

	c1, err := ports.CharacterRepository.Create(context.Background(), &model.Character{
		Name:      "Tester1",
		SessionID: &sessionID,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = ports.SessionRepository.Create(context.Background(), &model.Session{
		ID:    sessionID,
		Owner: c1,
	})
	if err != nil {
		t.Fatal(err)
	}

	c2, err := ports.CharacterRepository.Create(context.Background(), &model.Character{
		Name:      "Tester2",
		SessionID: &sessionID,
	})
	if err != nil {
		t.Fatal(err)
	}

	request := LookupRequest{
		SessionID: sessionID,
		OwnerID:   c1,
	}

	characters, err := Lookup(context.Background(), ports, request)
	if err != nil {
		t.Fatal(err)
	}

	if len(characters) != 2 {
		t.Fatal("expected 2 characters")
	}

	if characters[0].ID != c1 {
		t.Fatal("expected character in list but it isn't")
	}

	if characters[1].ID != c2 {
		t.Fatal("expected character in list but it isn't")
	}

}
