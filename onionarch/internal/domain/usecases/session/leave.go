package session

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	LeavePorts struct {
		SessionRepository   repository.SessionRepository
		CharacterRepository repository.CharacterRepository
	}

	LeaveRequest struct {
		SessionID   model.SessionID
		CharacterID int
	}
)

func Leave(ctx context.Context, ports LeavePorts, req LeaveRequest) error {
	_, err := ports.SessionRepository.FindByID(ctx, req.SessionID)
	if err != nil {
		return err
	}

	char, err := ports.CharacterRepository.FindByID(ctx, req.CharacterID)
	if err != nil {
		return err
	}

	char.SessionID = nil
	ports.CharacterRepository.Update(ctx, char)

	return nil
}
