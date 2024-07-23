package notification

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type (
	EventRecipient struct {
		SessionID   model.SessionID
		CharacterID int
	}

	EventSink struct {
		events map[EventRecipient][]model.Notification
	}
)

func NewEventSink() notification.Consumer {
	return &EventSink{
		events: make(map[EventRecipient][]model.Notification),
	}
}

func (es *EventSink) Consume(ctx context.Context, notificationChan <-chan model.Notification) error {
	for {
		select {
		case notification := <-notificationChan:
			recipient := EventRecipient{
				SessionID:   notification.SessionId,
				CharacterID: notification.FromId,
			}
			es.events[recipient] = append(es.events[recipient], notification)
			continue
		case <-ctx.Done():
			return nil
		}
	}
}

func (es *EventSink) CollectFor(recipient EventRecipient, offset int) []model.Notification {
	return es.events[recipient][offset:]
}
