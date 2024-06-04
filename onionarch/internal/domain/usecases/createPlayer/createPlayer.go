package createplayer

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	Ports struct {
		playerRepository    repository.PlayerRepository
		characterRepository repository.CharacterRepository
	}

	Request struct {
	}
)

func Create(ctx context.Context, ports Ports, reg Request) error {

	return nil
}
