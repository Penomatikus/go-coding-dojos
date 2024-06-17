package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	createcharacter "github.com/Penomatikus/onionarch/internal/domain/usecases/createCharacter"
)

type characterHandler struct {
	ctx                 context.Context
	characterRepository repository.CharacterRepository
	playerRepository    repository.PlayerRepository
}

func ProvideCharacterHandler(ctx context.Context, characterRepository repository.CharacterRepository) *characterHandler {
	return &characterHandler{
		ctx:                 ctx,
		characterRepository: characterRepository,
	}
}

// route: /api/v1/fatecore/character/new
func (handler *characterHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	handler.createCharacter(w, r)
}

func (handler *characterHandler) createCharacter(w http.ResponseWriter, r *http.Request) {
	if err := methodAllowed(http.MethodPost, w, r); err != nil {
		return
	}

	var request createcharacter.Request
	if err := decodeRequest(&request, w, r); err != nil {
		return
	}

	ports := createcharacter.Ports{
		PlayerRepository:    handler.playerRepository,
		CharacterRepository: handler.characterRepository,
	}

	if err := createcharacter.Create(handler.ctx, ports, request); err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
