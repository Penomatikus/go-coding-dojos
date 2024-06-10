package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	createplayer "github.com/Penomatikus/onionarch/internal/domain/usecases/createPlayer"
)

type characterHandler struct {
	ctx              context.Context
	playerRepository repository.PlayerRepository
}

func ProvideCharacterHandler(ctx context.Context, playerRepository repository.PlayerRepository) *characterHandler {
	return &characterHandler{
		ctx:              ctx,
		playerRepository: playerRepository,
	}
}

// route: /api/v1/fatecore/character/new
func (handler *characterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.createCharacter(w, r)
}

func (handler *characterHandler) createCharacter(w http.ResponseWriter, r *http.Request) {
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
