package http

import (
	"context"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type playerHandler struct {
	ctx              context.Context
	playerRepository repository.PlayerRepository
}

func ProvidePlayerHandler(ctx context.Context, playerRepository repository.PlayerRepository) *playerHandler {
	return &playerHandler{
		ctx:              ctx,
		playerRepository: playerRepository,
	}
}

// route: /api/fatecore/new/player
func (mgr *playerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	//TODO
}
