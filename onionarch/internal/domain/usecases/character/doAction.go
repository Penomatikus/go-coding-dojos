package character

import (
	"context"
	"errors"
	"math"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	DoPorts struct {
		CharacterRepository repository.CharacterRepository
		SessionRepository   repository.SessionRepository
	}

	DoRequest struct {
		ActionName  string
		CharacterID int
		Costs       int
		SessionID   model.SessionID
	}
)

func Do(ctx context.Context, ports DoPorts, request DoRequest) (int, error) {
	session, err := ports.SessionRepository.FindByID(ctx, request.SessionID)
	if err != nil {
		return 0, err
	}

	char, err := ports.CharacterRepository.FindByID(ctx, request.CharacterID)
	if err != nil {
		return 0, err
	}

	if char.SessionID == nil || string(*char.SessionID) != string(session.ID) {
		return 0, errors.New("character is not in that session")
	}

	if request.Costs < 0 {
		char.Points -= int(math.Abs(float64(request.Costs)))
	} else {
		char.Points += request.Costs
	}

	return char.Points, ports.CharacterRepository.Update(ctx, char)
}
