package joinsession

import (
	"context"
	"errors"

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

var ErrAnotherSession = errors.New("already in another session")

func Join(ctx context.Context, ports Ports, req Request) error {
	session, err := ports.SessionRepo.FindByID(ctx, req.SessionID())
	if err != nil {
		return err
	}

	char, err := ports.CharacterRepo.FindByID(ctx, req.CharacterID())
	if err != nil {
		return err
	}

	if char.SessionID != nil {
		return ErrAnotherSession
	}

	char.SessionID = &session.ID
	ports.CharacterRepo.Update(ctx, char)
	return nil
}
