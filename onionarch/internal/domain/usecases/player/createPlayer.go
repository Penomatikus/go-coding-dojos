package player

import (
	"context"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	CreatePorts struct {
		PlayerRepository repository.PlayerRepository
	}

	CreateRequest struct {
		Name string
	}
)

func Create(ctx context.Context, ports CreatePorts, req CreateRequest) error {
	return ports.PlayerRepository.Create(ctx, &model.Player{
		Name:      req.Name,
		CreatedAt: time.Now(),
	})
}
