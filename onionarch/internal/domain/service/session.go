package service

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

type Session interface {
	End(ctx context.Context, sessionID string) error
	Join(ctx context.Context, sessionID string, character model.Character) error
	Start(ctx context.Context, title string, ownerID int) error
}
