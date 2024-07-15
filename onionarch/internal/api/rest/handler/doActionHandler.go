package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
)

type (
	DoActionHandler struct {
		ctx     context.Context
		doPorts character.DoPorts
	}

	DoActionResponse struct {
		DiceRoll      int `json:"dice_roll"`
		CurrentPoints int `json:"current_points"`
	}
)

func NewDoActionHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	sessionRepository repository.SessionRepository) *DoActionHandler {
	return &DoActionHandler{
		ctx: ctx,
		doPorts: character.DoPorts{
			CharacterRepository: characterRepository,
			SessionRepository:   sessionRepository,
		},
	}
}

func (handler *DoActionHandler) DoAction(w http.ResponseWriter, r *http.Request) {
	handler.doAction(w, r)
}

func (handler *DoActionHandler) doAction(w http.ResponseWriter, r *http.Request) {
	var request character.DoRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	result := usecases.RollDice()
	request.Costs += result

	points, err := character.Do(handler.ctx, handler.doPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while creating chraracter: %v", err), http.StatusBadRequest)
		return
	}

	response := DoActionResponse{
		DiceRoll:      result,
		CurrentPoints: points,
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
