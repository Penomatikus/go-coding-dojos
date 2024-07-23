package session

import (
	"context"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
)

type (
	StartRequest struct {
		Title   string
		OwnerID int
	}

	StartPorts struct {
		CharacterRepository repository.CharacterRepository
		SessionRepository   repository.SessionRepository
		SessionIDGenerator  sessionid.Generator
	}
)

func Start(ctx context.Context, ports StartPorts, req StartRequest) (*model.SessionID, error) {
	sessionID, err := ports.SessionIDGenerator.GenerateSessionID()
	if err != nil {
		return nil, err
	}

	owner, err := ports.CharacterRepository.FindByID(ctx, req.OwnerID)
	if err != nil {
		return nil, err
	}

	err = ports.SessionRepository.Create(ctx, &model.Session{
		ID:        sessionID,
		CreatedAt: time.Now(),
		Title:     req.Title,
		Owner:     req.OwnerID,
	})
	if err != nil {
		return nil, err
	}

	owner.SessionID = &sessionID
	return &sessionID, ports.CharacterRepository.Update(ctx, owner)
}
