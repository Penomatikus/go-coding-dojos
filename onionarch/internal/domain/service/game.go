package service

import "context"

type Game interface {
	EndSession(ctx context.Context) error
	JoinSession(ctx context.Context) error
	SessionMetrics(ctx context.Context) error
	StartSession(ctx context.Context) error
}
