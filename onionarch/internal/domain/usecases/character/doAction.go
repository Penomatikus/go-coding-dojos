package character

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type (
	DoPorts struct {
		CharacterRepository   repository.CharacterRepository
		SessionRepository     repository.SessionRepository
		NotificationPublisher notification.Publisher
	}

	DoRequest struct {
		ActionName  string
		CharacterID int
		Costs       int
		SessionID   model.SessionID
	}
)

func Do(ctx context.Context, ports DoPorts, request DoRequest) (int, error) {
	session, err := ports.SessionRepository.FindByID(ctx, request.SessionID)
	if err != nil {
		return 0, err
	}

	char, err := ports.CharacterRepository.FindByID(ctx, request.CharacterID)
	if err != nil {
		return 0, err
	}

	if char.SessionID == nil || string(*char.SessionID) != string(session.ID) {
		return 0, errors.New("character is not in that session")
	}

	if request.Costs < 0 {
		char.Points -= int(math.Abs(float64(request.Costs)))
	} else {
		char.Points += request.Costs
	}

	err = ports.NotificationPublisher.Publish(ctx, model.Notification{
		CreatedAt: time.Now(),
		SessionId: request.SessionID,
		FromId:    char.ID,
		Body:      []byte(fmt.Sprintf("%d points spent", request.Costs)),
	})
	if err != nil {
		return 0, err
	}

	return char.Points, ports.CharacterRepository.Update(ctx, char)
}
