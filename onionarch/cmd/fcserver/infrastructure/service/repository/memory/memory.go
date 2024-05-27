package memory

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/service"
)

type sessionRepository struct {
	sessions map[string]model.Session
}

var _ service.SessionRepository = &sessionRepository{}

func (mr *sessionRepository) LoadSession(ctx context.Context, sessionID string) (model.Session, error) {
	s, ok := mr.sessions[sessionID]
	if !ok {
		return model.Session{}, service.ErrNotFound
	}
	return s, nil
}

func (mr *sessionRepository) CreateSession(ctx context.Context, session model.Session) error {
	if _, ok := mr.sessions[session.SessionID]; ok {
		return service.ErrAlreadyExists
	}

	mr.sessions[session.SessionID] = session
	return nil
}

func ProvideRepository() service.SessionRepository {
	return &sessionRepository{}
}
