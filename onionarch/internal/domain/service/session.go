package service

import (
	"context"
)

type Session interface {
	End(ctx context.Context, sessionID string) error
	Join(ctx context.Context, sessionID string) error
	New(ctx context.Context) (sessionID string, err error)
}
