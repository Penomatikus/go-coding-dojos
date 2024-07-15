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
		Title string
		Owner int
	}

	StartPorts struct {
		SessionRepository  repository.SessionRepository
		SessionIDGenerator sessionid.Generator
	}
)

func Start(ctx context.Context, ports StartPorts, req StartRequest) (*model.SessionID, error) {
	sessionID, err := ports.SessionIDGenerator.GenerateSessionID()
	if err != nil {
		return nil, err
	}

	return &sessionID, ports.SessionRepository.Create(ctx, &model.Session{
		ID:        sessionID,
		CreatedAt: time.Now(),
		Title:     req.Title,
		Owner:     req.Owner,
	})
}
