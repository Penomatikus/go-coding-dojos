package notification

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type (
	// Just a convinient type for chan<- model.Notification
	EventSubscriber chan model.Notification

	EventBus struct {
		// Note: We might just have one EventSubscriber in this kata, but its cool to have the opportunity to have more
		sessionSubscriber map[model.SessionID][]EventSubscriber
	}
)

func NewEventBus() notification.Publisher {
	return &EventBus{
		sessionSubscriber: make(map[model.SessionID][]EventSubscriber, 0),
	}
}

func (eb *EventBus) Publish(ctx context.Context, notification model.Notification) error {
	for _, subscriber := range eb.sessionSubscriber[notification.SessionId] {
		subscriber <- notification
	}
	return nil
}

// Subscribe with a subscriber for a specific recipient within the subscriber
func (eb *EventBus) Subscribe(subscriber EventSubscriber, sessionID model.SessionID) {
	eb.sessionSubscriber[sessionID] = append(eb.sessionSubscriber[sessionID], subscriber)
}
