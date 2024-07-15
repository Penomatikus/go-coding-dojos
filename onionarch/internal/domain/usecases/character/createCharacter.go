package character

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	CreatePorts struct {
		CharacterRepository repository.CharacterRepository
	}

	CreateRequest struct {
		Name, Description string
	}
)

func Create(ctx context.Context, ports CreatePorts, reg CreateRequest) error {
	return ports.CharacterRepository.Create(ctx, &model.Character{
		Name:        reg.Name,
		Description: reg.Description,
	})
}
