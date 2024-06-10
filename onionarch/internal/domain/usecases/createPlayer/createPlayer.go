package createplayer

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type Ports struct {
	PlayerRepository repository.PlayerRepository
}

func Create(ctx context.Context, ports Ports, playerName string) error {
	return ports.PlayerRepository.Create(ctx, &model.Player{
		Name: playerName,
	})
}
