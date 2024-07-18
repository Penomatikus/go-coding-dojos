package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session"
	infraNotification "github.com/Penomatikus/onionarch/internal/infrastructure/notification"
)

type SessionHandler struct {
	ctx                    context.Context
	notificationPublisher  *infraNotification.EventBus
	notificationSubscriber infraNotification.EventSubscriber
	startsessionPorts      session.StartPorts
	joinsessionPorts       session.JoinPorts
	leavesessionPorts      session.LeavePorts
}

func NewSessionHandler(ctx context.Context,
	notificationPublisher notification.Publisher,
	notificationSubscriber infraNotification.EventSubscriber,
	characterRepository repository.CharacterRepository,
	sessionIDGen sessionid.Generator,
	sessionRepository repository.SessionRepository,
) *SessionHandler {

	eventbus, ok := notificationPublisher.(*infraNotification.EventBus)
	if !ok {
		panic("failed to cast interface ")
	}

	return &SessionHandler{
		ctx:                    ctx,
		notificationPublisher:  eventbus,
		notificationSubscriber: notificationSubscriber,
		startsessionPorts: session.StartPorts{
			SessionRepository:  sessionRepository,
			SessionIDGenerator: sessionIDGen,
		},
		joinsessionPorts: session.JoinPorts{
			SessionRepository:   sessionRepository,
			CharacterRepository: characterRepository,
		},
		leavesessionPorts: session.LeavePorts{
			SessionRepository:   sessionRepository,
			CharacterRepository: characterRepository,
		},
	}
}

// route: /api/v1/fatecore/session/new
func (handler *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	handler.startSession(w, r)
}

// route: /api/v1/fatecore/session/{sessionid}/join
func (handler *SessionHandler) JoinSession(w http.ResponseWriter, r *http.Request) {
	handler.joinSession(w, r)
}

// route: /api/v1/fatecore/session/{sessionid}/leave
func (handler *SessionHandler) LeaveSession(w http.ResponseWriter, r *http.Request) {
	handler.leaveSession(w, r)
}

func (handler *SessionHandler) startSession(w http.ResponseWriter, r *http.Request) {
	var request session.StartRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	id, err := session.Start(handler.ctx, handler.startsessionPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating new session: %v", err), http.StatusBadRequest)
		return
	}

	handler.notificationPublisher.Subscribe(handler.notificationSubscriber)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("application", "plain/text")
	fmt.Fprint(w, *id)
}

func (handler *SessionHandler) joinSession(w http.ResponseWriter, r *http.Request) {
	var request session.JoinRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	sID, ok := rest.PathValues(r, "sessionid")["sessionid"]
	if !ok {
		http.Error(w, "error while reading session id from path", http.StatusBadRequest)
		return
	}
	request.SessionID = model.SessionID(sID)

	err := session.Join(handler.ctx, handler.joinsessionPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *SessionHandler) leaveSession(w http.ResponseWriter, r *http.Request) {
	var request session.LeaveRequest
	if err := rest.DecodeRequest(&request, w, r); err != nil {
		return
	}

	sID, ok := rest.PathValues(r, "sessionid")["sessionid"]
	if !ok {
		http.Error(w, "error while reading session id from path", http.StatusBadRequest)
		return
	}
	request.SessionID = model.SessionID(sID)

	err := session.Leave(handler.ctx, handler.leavesessionPorts, request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error leaving session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
