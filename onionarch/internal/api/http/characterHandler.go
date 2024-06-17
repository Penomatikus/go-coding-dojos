package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	createcharacter "github.com/Penomatikus/onionarch/internal/domain/usecases/character/createCharacter"
	updatecharacter "github.com/Penomatikus/onionarch/internal/domain/usecases/character/updateCharacter"
)

type characterHandler struct {
	ctx         context.Context
	createPorts createcharacter.Ports
	updatePorts updatecharacter.Ports
}

func ProvideCharacterHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	playerRepository repository.PlayerRepository) *characterHandler {
	return &characterHandler{
		ctx: ctx,
		createPorts: createcharacter.Ports{
			PlayerRepository:    playerRepository,
			CharacterRepository: characterRepository,
		},
		updatePorts: updatecharacter.Ports{
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

	var request createcharacter.Request
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	if err := createcharacter.Create(handler.ctx, handler.createPorts, request); err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *characterHandler) updateCharacter(w http.ResponseWriter, r *http.Request) {
	if err := methodAllowed(http.MethodPost, w, r); err != nil {
		return
	}

	var request updatecharacter.Request
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	if err := updatecharacter.Update(handler.ctx, handler.updatePorts, request); err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
