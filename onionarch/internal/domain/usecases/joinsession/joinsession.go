package joinsession

import (
	"context"
	"errors"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases"
)

type (
	Ports struct {
		SessionRepo   repository.SessionRepository
		CharacterRepo repository.CharacterRepository
	}

	Request struct {
		SessionID   usecases.SessionID
		CharacterID usecases.CharacterID
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
