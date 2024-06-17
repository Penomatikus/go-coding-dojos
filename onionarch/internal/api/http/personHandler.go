package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	createplayer "github.com/Penomatikus/onionarch/internal/domain/usecases/player/createPlayer"
)

type playerHandler struct {
	ctx               context.Context
	createplayerPorts createplayer.Ports
}

func ProvidePlayerHandler(ctx context.Context, playerRepository repository.PlayerRepository) *playerHandler {
	return &playerHandler{
		ctx: ctx,
		createplayerPorts: createplayer.Ports{
			PlayerRepository: playerRepository,
		},
	}
}

// route: /api/v1/fatecore/player/new
func (handler *playerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	if methodAllowed(http.MethodPost, w, r) != nil {
		return
	}

	var request createplayer.Request
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	err := createplayer.Create(handler.ctx, handler.createplayerPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating player: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
