package notification

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type (
	// EventSubscriber defines a technical subscriber channel for Notifications
	//
	// Note: Just a convinient type for chan<- model.Notification
	EventSubscriber chan model.Notification

	EventBus struct {
		// Note: We might just have one EventSubscriber in this kata, but its cool to have the opportunity to have more
		subscribers []EventSubscriber
	}
)

func NewEventBus() notification.Publisher {
	return &EventBus{
		subscribers: make([]EventSubscriber, 0),
	}
}

func (eb *EventBus) Publish(ctx context.Context, notification model.Notification) error {
	for _, subscriber := range eb.subscribers {
		subscriber <- notification
	}
	return nil
}

// Subscribe with a subscriber for a specific recipient within the subscriber
func (at *EventBus) Subscribe(subscriber EventSubscriber) {
	at.subscribers = append(at.subscribers, subscriber)
}
