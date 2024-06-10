package createcharacter

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	Ports struct {
		playerRepository    repository.PlayerRepository
		characterRepository repository.CharacterRepository
	}

	Request struct {
		PlayerID          int
		Name, Description string
	}
)

func Create(ctx context.Context, ports Ports, reg Request) error {
	_, err := ports.playerRepository.FindByID(ctx, reg.PlayerID)
	if err != nil {
		return err
	}

	return ports.characterRepository.Create(ctx, &model.Character{
		Name:        reg.Name,
		Description: reg.Description,
		PlayerID:    reg.PlayerID,
	})
}
