package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
)

type CharacterHandler struct {
	ctx         context.Context
	createPorts character.CreatePorts
	updatePorts character.UpdatePorts
}

func NewCharacterHandler(ctx context.Context, characterRepository repository.CharacterRepository) *CharacterHandler {
	return &CharacterHandler{
		ctx: ctx,
		createPorts: character.CreatePorts{
			CharacterRepository: characterRepository,
		},
		updatePorts: character.UpdatePorts{
			CharacterRepository: characterRepository,
		},
	}
}

// route: /api/v1/fatecore/character/new
func (handler *CharacterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.createCharacter(w, r)
}

// route: /api/v1/fatecore/character/{id}/update
func (handler *CharacterHandler) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.updateCharacter(w, r)
}

func (handler *CharacterHandler) createCharacter(w http.ResponseWriter, r *http.Request) {
	var request character.CreateRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}
	charID, err := character.Create(handler.ctx, handler.createPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while creating chraracter: %v", err), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(struct{ CharacterID int }{CharacterID: charID})
	if err != nil {
		http.Error(w, "error while serializing norifications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (handler *CharacterHandler) updateCharacter(w http.ResponseWriter, r *http.Request) {
	var request character.UpdateRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	if err := character.Update(handler.ctx, handler.updatePorts, request); err != nil {
		http.Error(w, fmt.Sprintf("error while updating chraracter: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
