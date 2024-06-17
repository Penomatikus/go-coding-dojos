package updatecharacter

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	Ports struct {
		PlayerRepository    repository.PlayerRepository
		CharacterRepository repository.CharacterRepository
	}

	Request struct {
		ID        int
		PlayerID  int
		Points    *int
		SessionID *model.SessionID
	}
)

func Update(ctx context.Context, ports Ports, req Request) error {
	if req.Points == nil && req.SessionID == nil {
		return nil
	}

	_, err := ports.PlayerRepository.FindByID(ctx, req.PlayerID)
	if err != nil {
		return err
	}

	_, err = ports.CharacterRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	update := &model.Character{
		ID: req.ID,
	}

	if req.Points != nil {
		update.Points = *req.Points
	}

	if req.SessionID != nil {
		update.SessionID = req.SessionID
	}

	return ports.CharacterRepository.Update(ctx, update)
}
