package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/character"
)

type (
	DoHandler struct {
		ctx     context.Context
		doPorts character.DoPorts
	}

	DoResponse struct {
		DiceRoll      int `json:"dice_roll"`
		CurrentPoints int `json:"current_points"`
	}
)

func NewDoHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	sessionRepository repository.SessionRepository,
	notificationPublisher notification.Publisher,
) *DoHandler {
	return &DoHandler{
		ctx: ctx,
		doPorts: character.DoPorts{
			CharacterRepository:   characterRepository,
			SessionRepository:     sessionRepository,
			NotificationPublisher: notificationPublisher,
		},
	}
}

func (handler *DoHandler) DoAction(w http.ResponseWriter, r *http.Request) {
	handler.doAction(w, r)
}

func (handler *DoHandler) DoPoints(w http.ResponseWriter, r *http.Request) {
	handler.doPoints(w, r)
}

func (handler *DoHandler) doAction(w http.ResponseWriter, r *http.Request) {
	var request character.DoRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	result := usecases.RollDice()
	request.Costs += result

	points, err := character.Do(handler.ctx, handler.doPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while doing action: %v", err), http.StatusInternalServerError)
		return
	}

	response := DoResponse{
		DiceRoll:      result,
		CurrentPoints: points,
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (handler *DoHandler) doPoints(w http.ResponseWriter, r *http.Request) {
	var request character.DoRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	points, err := character.Do(handler.ctx, handler.doPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while doing points: %v", err), http.StatusInternalServerError)
		return
	}

	response := DoResponse{
		DiceRoll:      0,
		CurrentPoints: points,
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
