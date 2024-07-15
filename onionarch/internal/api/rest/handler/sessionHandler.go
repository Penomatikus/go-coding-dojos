package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/api/rest"
	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/session"
)

type SessionHandler struct {
	ctx               context.Context
	startsessionPorts session.StartPorts
	joinsessionPorts  session.JoinPorts
	leavesessionPorts session.LeavePorts
}

func NewSessionHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	sessionIDGen sessionid.Generator,
	sessionRepository repository.SessionRepository,
) *SessionHandler {
	return &SessionHandler{
		ctx: ctx,
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
