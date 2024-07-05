package session

import (
	"context"
	"errors"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	JoinPorts struct {
		SessionRepository   repository.SessionRepository
		CharacterRepository repository.CharacterRepository
	}

	JoinRequest struct {
		SessionID   model.SessionID
		CharacterID int
	}
)

var ErrAnotherSession = errors.New("already in another session")

func Join(ctx context.Context, ports JoinPorts, req JoinRequest) error {
	session, err := ports.SessionRepository.FindByID(ctx, req.SessionID)
	if err != nil {
		return err
	}

	char, err := ports.CharacterRepository.FindByID(ctx, req.CharacterID)
	if err != nil {
		return err
	}

	if char.SessionID != nil {
		return ErrAnotherSession
	}

	char.SessionID = &session.ID
	ports.CharacterRepository.Update(ctx, char)
	return nil
}
