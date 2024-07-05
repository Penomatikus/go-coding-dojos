package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
)

type characterHandler struct {
	ctx         context.Context
	createPorts character.CreatePorts
	updatePorts character.UpdatePorts
}

func ProvideCharacterHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	playerRepository repository.PlayerRepository) *characterHandler {
	return &characterHandler{
		ctx: ctx,
		createPorts: character.CreatePorts{
			PlayerRepository:    playerRepository,
			CharacterRepository: characterRepository,
		},
		updatePorts: character.UpdatePorts{
			PlayerRepository:    playerRepository,
			CharacterRepository: characterRepository,
		},
	}
}

// route: /api/v1/fatecore/character/new
func (handler *characterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.createCharacter(w, r)
}

// route: /api/v1/fatecore/character/{id}/update
func (handler *characterHandler) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.updateCharacter(w, r)
}

func (handler *characterHandler) createCharacter(w http.ResponseWriter, r *http.Request) {
	if err := methodAllowed(http.MethodPost, w, r); err != nil {
		return
	}

	var request character.CreateRequest
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	if err := character.Create(handler.ctx, handler.createPorts, request); err != nil {
		http.Error(w, fmt.Sprintf("error while creating chraracter: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *characterHandler) updateCharacter(w http.ResponseWriter, r *http.Request) {
	if err := methodAllowed(http.MethodPost, w, r); err != nil {
		return
	}

	var request character.UpdateRequest
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	if err := character.Update(handler.ctx, handler.updatePorts, request); err != nil {
		http.Error(w, fmt.Sprintf("error while updating chraracter: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
