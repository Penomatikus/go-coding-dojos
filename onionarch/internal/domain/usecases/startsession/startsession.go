package startsession

import (
	"context"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	SessionID func() model.SessionID
	Request   struct {
		ID      SessionID
		Title   string
		OwnerID int
	}

	Ports struct {
		SessionRepo repository.SessionRepository
	}
)

func Start(ctx context.Context, ports Ports, req Request) (*model.SessionID, error) {
	sessionID := req.ID()
	return &sessionID, ports.SessionRepo.Create(ctx, &model.Session{
		ID:        sessionID,
		CreatedAt: time.Now(),
		Title:     req.Title,
		Owner:     req.OwnerID,
	})
}
