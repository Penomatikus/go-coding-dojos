package handler

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	infraNotification "github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

func Test_ReceiveNotification(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sink := infraNotification.NewEventSink()
	notificationHandler := NewNotificationHandler(
		ctx,
		sink,
	)

	subscriberChan := make(infraNotification.EventSubscriber)
	notificationPublisher := infraNotification.NewEventBus()
	eventBus, ok := notificationPublisher.(*infraNotification.EventBus)
	if !ok {
		panic("failed to cast interface ")
	}

	sessionID := model.SessionID("1337")
	eventBus.Subscribe(subscriberChan, sessionID)

	go sink.Consume(ctx, subscriberChan)
	defer cancel()

	eventBus.Publish(ctx, model.Notification{
		Body:      "Hallo Welt",
		CreatedAt: time.Now(),
		FromId:    1,
		SessionId: sessionID,
	})

	req := httptest.NewRequest("GET", "/api/v1/fatecore/session/1337/notification?charID=1&offset=0", nil)
	rec := httptest.NewRecorder()

	req.SetPathValue("sessionid", "1337")

	notificationHandler.CollectNotification(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if len(data) == 0 {
		t.Fatal("no data in body")
	}
}
