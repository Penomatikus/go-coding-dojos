package createplayer

import (
	"context"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	Ports struct {
		PlayerRepository repository.PlayerRepository
	}

	Request struct {
		Name string
	}
)

func Create(ctx context.Context, ports Ports, req Request) error {
	return ports.PlayerRepository.Create(ctx, &model.Player{
		Name:      req.Name,
		CreatedAt: time.Now(),
	})
}
