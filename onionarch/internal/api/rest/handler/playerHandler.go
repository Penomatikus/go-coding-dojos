package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/player"
)

type PlayerHandler struct {
	ctx               context.Context
	createplayerPorts player.CreatePorts
}

func NewPlayerHandler(ctx context.Context, playerRepository repository.PlayerRepository) *PlayerHandler {
	return &PlayerHandler{
		ctx: ctx,
		createplayerPorts: player.CreatePorts{
			PlayerRepository: playerRepository,
		},
	}
}

// route: /api/v1/fatecore/player/new
func (handler *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var request player.CreateRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	err := player.Create(handler.ctx, handler.createplayerPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating player: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
