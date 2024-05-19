package service

import (
	"context"
)

type Notification interface {
	Notifiy(ctx context.Context, data []byte)
}
