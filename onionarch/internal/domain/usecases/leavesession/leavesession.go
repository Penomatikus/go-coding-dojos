package leavesession

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	Ports struct {
		SessionRepo   repository.SessionRepository
		CharacterRepo repository.CharacterRepository
	}

	SessionID   func() model.SessionID
	CharacterID func() model.CharacterID

	Request struct {
		SessionID   SessionID
		CharacterID CharacterID
	}
)

func Leave(ctx context.Context, ports Ports, req Request) error {
	_, err := ports.SessionRepo.FindByID(ctx, req.SessionID())
	if err != nil {
		return err
	}

	char, err := ports.CharacterRepo.FindByID(ctx, req.CharacterID())
	if err != nil {
		return err
	}

	char.SessionID = nil
	ports.CharacterRepo.Update(ctx, char)

	return nil
}
