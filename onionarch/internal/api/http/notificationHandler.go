package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	infraNotification "github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

type notificationHandler struct {
	ctx     context.Context
	sink    map[model.SessionID][]model.Notification
	service notification.Service
}

func ProvideNotificationHandler(ctx context.Context, Service notification.Service) *notificationHandler {
	notificationSink := make(map[model.SessionID][]model.Notification)

	return &notificationHandler{
		ctx:     ctx,
		sink:    notificationSink,
		service: Service,
	}
}

// route: /api/v1/fatecore/session/{sessionid}/notification POST
func (handler *notificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	handler.sendNotification(w, r)
}

// // route: /api/v1/fatecore/session/{sessionid}/notification GET
func (handler *notificationHandler) ReceiveNotification(w http.ResponseWriter, r *http.Request) {
	handler.receiveNotification(w, r)
}

func (handler *notificationHandler) sendNotification(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
		http.Error(w, fmt.Sprintf("error seinding notification to session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *notificationHandler) receiveNotification(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodGet, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.PathValue("sessionid")
	if len(sessionID) == 0 {
		http.Error(w, "error while reading session id from path", http.StatusInternalServerError)
		return
	}

	var request struct {
		Offset int
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	notificationReader := infraNotification.JSONSinkReader{
		Offset:    request.Offset,
		SessionID: model.SessionID(sessionID),
		Sink:      &handler.sink,
	}

	responsebody := []byte{}
	err := handler.service.Read(handler.ctx, responsebody, notificationReader)
	if err != nil {
		http.Error(w, fmt.Sprintf("error leaving session: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responsebody)
}
