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
		PlayerRepository   repository.PlayerRepository
		SessionRepository  repository.SessionRepository
		SessionIDGenerator sessionid.Generator
	}
)

func Start(ctx context.Context, ports StartPorts, req StartRequest) (*model.SessionID, error) {
	_, err := ports.PlayerRepository.FindByID(ctx, req.Owner)
	if err != nil {
		return nil, err
	}

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
