package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/notification"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/joinsession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/leavesession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/startsession"
)

type sessionHandler struct {
	ctx                 context.Context
	characterRepository repository.CharacterRepository
	notificationService notification.Service
	sessionIDGen        sessionid.Generator
	sessionRepository   repository.SessionRepository
}

func ProvidesessionHandler(ctx context.Context,
	characterRepository repository.CharacterRepository,
	notificationService notification.Service,
	sessionIDGen sessionid.Generator,
	sessionRepository repository.SessionRepository,
) *sessionHandler {
	return &sessionHandler{
		ctx:                 ctx,
		characterRepository: characterRepository,
		notificationService: notificationService,
		sessionIDGen:        sessionIDGen,
		sessionRepository:   sessionRepository,
	}
}

// route: /api/v1/fatecore/session/new
func (handler *sessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	handler.startSession(w, r)
}

// route: /api/v1/fatecore/session/{sessionid}/join
func (handler *sessionHandler) JoinSession(w http.ResponseWriter, r *http.Request) {
	handler.joinSession(w, r)
}

// route: /api/v1/fatecore/session/{sessionid}/leave
func (handler *sessionHandler) LeaveSession(w http.ResponseWriter, r *http.Request) {
	handler.leaveSession(w, r)
}

func (handler *sessionHandler) startSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Title   string `json:"title"`
		OwnerID int    `json:"owner_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	startsessionPorts := startsession.Ports{
		SessionRepository:  handler.sessionRepository,
		SessionIDGenerator: handler.sessionIDGen,
	}

	id, err := startsession.Start(handler.ctx, startsessionPorts, startsession.Request{
		Title:   req.Title,
		OwnerID: req.OwnerID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("error creating new session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("application", "plain/text")
	fmt.Fprint(w, id)
}

func (handler *sessionHandler) joinSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var characterID int
	if err := json.NewDecoder(r.Body).Decode(&characterID); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	joinSessionPorts := joinsession.Ports{
		SessionRepository:   handler.sessionRepository,
		CharacterRepository: handler.characterRepository,
	}

	sessionID := r.PathValue("sessionid")
	if len(sessionID) == 0 {
		http.Error(w, "error while reading session id from path", http.StatusInternalServerError)
		return
	}

	err := joinsession.Join(handler.ctx, joinSessionPorts, joinsession.Request{
		SessionID:   model.SessionID(sessionID),
		CharacterID: characterID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *sessionHandler) leaveSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var characterID int
	if err := json.NewDecoder(r.Body).Decode(&characterID); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	leaveSessionPorts := leavesession.Ports{
		SessionRepository:   handler.sessionRepository,
		CharacterRepository: handler.characterRepository,
	}

	sessionID := r.PathValue("sessionid")
	if len(sessionID) == 0 {
		http.Error(w, "error while reading session id from path", http.StatusInternalServerError)
		return
	}

	err := leavesession.Leave(handler.ctx, leaveSessionPorts, leavesession.Request{
		SessionID:   model.SessionID(sessionID),
		CharacterID: characterID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("error leaving session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
