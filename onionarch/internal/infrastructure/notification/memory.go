package notification

import (
	"context"
	"encoding/json"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
)

type SinkService struct {
	Notifications map[model.SessionID][]model.Notification
}

func ProvideNotificationService() notification.Service {
	return &SinkService{
		make(map[model.SessionID][]model.Notification),
	}
}

func (s *SinkService) Send(ctx context.Context, data []byte) error {
	var note model.Notification
	err := json.Unmarshal(data, &note)
	if err != nil {
		return err
	}
	s.Notifications[note.SessionId] = append(s.Notifications[note.SessionId], note)
	return nil
}

func (s *SinkService) Receive(ctx context.Context, id model.SessionID, offset int) ([]model.Notification, error) {
	return s.Notifications[id][offset:], nil
}
