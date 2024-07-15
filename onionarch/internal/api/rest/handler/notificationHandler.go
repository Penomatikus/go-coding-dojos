package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	infraNotification "github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

type NotificationHandler struct {
	ctx     context.Context
	sink    map[model.SessionID][]model.Notification
	service notification.Service
}

func NewNotificationHandler(ctx context.Context, Service notification.Service) *NotificationHandler {
	notificationSink := make(map[model.SessionID][]model.Notification)

	return &NotificationHandler{
		ctx:     ctx,
		sink:    notificationSink,
		service: Service,
	}
}

// route: /api/v1/fatecore/session/{sessionid}/notification POST
func (handler *NotificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	handler.sendNotification(w, r)
}

// // route: /api/v1/fatecore/session/{sessionid}/notification GET
func (handler *NotificationHandler) CollectNotification(w http.ResponseWriter, r *http.Request) {
	handler.collectNotification(w, r)
}

func (handler *NotificationHandler) sendNotification(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading body from request: %v", err), http.StatusInternalServerError)
		return
	}

	notificationWriter := infraNotification.JSONSinkWriter{
		Sink: &handler.sink,
	}

	err = handler.service.Send(handler.ctx, body, notificationWriter)
	if err != nil {
		http.Error(w, fmt.Sprintf("error sending notification to session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *NotificationHandler) collectNotification(w http.ResponseWriter, r *http.Request) {
	sID, ok := rest.PathValues(r, "sessionid")["sessionid"]
	if !ok {
		http.Error(w, "error while reading session id from path", http.StatusBadRequest)
		return
	}

	var request struct{ Offset int }
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	notificationReader := infraNotification.JSONSinkReader{
		Notifications: handler.sink[model.SessionID(sID)][request.Offset:],
	}

	out, err := handler.service.Read(handler.ctx, &notificationReader)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while reading session notifications: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
