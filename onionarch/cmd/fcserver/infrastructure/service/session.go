package service

import (
	"context"
	"fmt"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/service"
	"github.com/google/uuid"
)

var _ service.Session = &gameSessionService{}

type gameSessionService struct {
	repository service.SessionRepository
}

func Provide(r service.SessionRepository) service.Session {
	return &gameSessionService{
		repository: r,
	}
}

func (s *gameSessionService) End(ctx context.Context, sessionID string) error {
	return nil
}

func (s *gameSessionService) Join(ctx context.Context, sessionID string) error {
	return nil
}

func (s *gameSessionService) New(ctx context.Context) (sessionID string, err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("failed to create session ID: %s", err)
	}

	session := model.Session{
		SessionID: id.String(),
		Game: model.Game{
			Players: []model.Player{},
		},
	}

	return id.String(), s.repository.CreateSession(ctx, session)
}
