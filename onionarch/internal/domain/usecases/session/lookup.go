package session

import (
	"context"
	"errors"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	LookupRequest struct {
		SessionID model.SessionID
		OwnerID   int
	}

	LookupPorts struct {
		SessionRepository   repository.SessionRepository
		CharacterRepository repository.CharacterRepository
	}
)

func Lookup(ctx context.Context, ports LookupPorts, req LookupRequest) ([]model.Character, error) {
	session, err := ports.SessionRepository.FindByID(ctx, req.SessionID)
	if err != nil {
		return nil, err
	}

	_, err = ports.CharacterRepository.FindByID(ctx, req.OwnerID)
	if err != nil {
		return nil, err
	}

	if session.Owner != req.OwnerID {
		return nil, errors.New("only session owners can lookup members")
	}

	return ports.CharacterRepository.FindBySession(ctx, req.SessionID)
}
