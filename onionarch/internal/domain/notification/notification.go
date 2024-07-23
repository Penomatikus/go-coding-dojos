package notification

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

type (
	Publisher interface {
		Publish(ctx context.Context, notification model.Notification) error
	}

	Consumer interface {
		Consume(ctx context.Context, notificationChan <-chan model.Notification) error
	}
)
