package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

func Test_SendNotification(t *testing.T) {
	ctx := context.Background()
	notificationHandler := ProvideNotificationHandler(
		ctx,
		notification.PrivideService(),
	)

	jsonData, err := json.Marshal(model.Notification{
		SessionId: "1337",
		FromId:    1,
		Body:      "Hello Moto!",
	})

	if err != nil {
		t.Fatalf("%s: Error marshalling data to JSON:", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/fatecore/session/1337/notification", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()

	notificationHandler.SendNotification(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}
}

func Test_ReceiveNotification(t *testing.T) {
	ctx := context.Background()
	notificationHandler := ProvideNotificationHandler(
		ctx,
		notification.PrivideService(),
	)

	jsonData, err := json.Marshal(model.Notification{
		SessionId: "1337",
		FromId:    1,
		Body:      "Hello Moto!!!",
	})

	if err != nil {
		t.Fatalf("%s: Error marshalling data to JSON:", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/fatecore/session/1337/notification", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()

	notificationHandler.SendNotification(rec, req)
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", res.StatusCode)
	}

	jsonData, err = json.Marshal(struct{ Offset int }{Offset: 0})
	if err != nil {
		fmt.Println("Error marshalling data to JSON:", err)
		return
	}

	req = httptest.NewRequest("GET", "/api/v1/fatecore/session/1337/notification", bytes.NewReader(jsonData))
	rec = httptest.NewRecorder()

	notificationHandler.CollectNotification(rec, req)
	res = rec.Result()
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
