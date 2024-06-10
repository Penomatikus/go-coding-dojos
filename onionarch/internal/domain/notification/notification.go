package notification

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

type Service interface {
	Send(ctx context.Context, data []byte) error
	Receive(ctx context.Context, id model.SessionID, offset int) ([]model.Notification, error)
}
