package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	infraNotification "github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

type NotificationHandler struct {
	ctx      context.Context
	sink     map[model.SessionID][]model.Notification
	consumer notification.Consumer
}

func NewNotificationHandler(ctx context.Context, consumer notification.Consumer) *NotificationHandler {
	notificationSink := make(map[model.SessionID][]model.Notification)

	return &NotificationHandler{
		ctx:      ctx,
		sink:     notificationSink,
		consumer: consumer,
	}
}

// // route: /api/v1/fatecore/session/{sessionid}/notification GET
func (handler *NotificationHandler) CollectNotification(w http.ResponseWriter, r *http.Request) {
	handler.collectNotification(w, r)
}

func (handler *NotificationHandler) collectNotification(w http.ResponseWriter, r *http.Request) {
	sID, ok := rest.PathValues(r, "sessionid")["sessionid"]
	if !ok {
		http.Error(w, "error while reading session id from path", http.StatusBadRequest)
		return
	}

	var request struct{ CharID, Offset int }
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	sink, ok := handler.consumer.(*infraNotification.EventSink)
	if !ok {
		panic("interface to type cast failed")
	}

	notificationForChat := sink.CollectFor(
		infraNotification.EventRecipient{
			SessionID:   model.SessionID(sID),
			CharacterID: request.CharID,
		}, request.Offset)

	out, err := json.Marshal(notificationForChat)
	if err != nil {
		http.Error(w, "error while serializing norifications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
