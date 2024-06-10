package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	createplayer "github.com/Penomatikus/onionarch/internal/domain/usecases/createPlayer"
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

// route: /api/v1/fatecore/player/new
func (handler *playerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var playerName string
	if err := json.NewDecoder(r.Body).Decode(&playerName); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	err := createplayer.Create(handler.ctx, createplayer.Ports{PlayerRepository: handler.playerRepository}, playerName)
	if err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
