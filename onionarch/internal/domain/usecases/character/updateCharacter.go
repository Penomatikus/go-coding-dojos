package character

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	UpdatePorts struct {
		CharacterRepository repository.CharacterRepository
	}

	UpdateRequest struct {
		ID        int
		Points    *int
		SessionID *model.SessionID
	}
)

func Update(ctx context.Context, ports UpdatePorts, req UpdateRequest) error {
	if req.Points == nil {
		return nil
	}
	_, err := ports.CharacterRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	update := &model.Character{
		ID:        req.ID,
		Points:    *req.Points,
		SessionID: req.SessionID,
	}
	return ports.CharacterRepository.Update(ctx, update)
}
