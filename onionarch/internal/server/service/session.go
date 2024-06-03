package service

import (
	"context"
	"fmt"
	"time"

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

func (s *gameSessionService) Join(ctx context.Context, sessionID string, character model.Character) error {
	return nil
}

func (s *gameSessionService) Start(ctx context.Context, title string, ownerID int) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("failed to create session ID: %s", err)
	}

	return s.repository.CreateSession(ctx, &model.Session{
		ID:         id.String(),
		Characters: make([]model.Character, 0),
		CreatedAt:  time.Now(),
		Title:      title,
	})
}
