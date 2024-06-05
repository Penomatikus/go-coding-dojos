package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/joinsession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/leavesession"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/startsession"
)

type SessionHandler struct {
	ctx                 context.Context
	sessionRepository   repository.SessionRepository
	characterRepository repository.CharacterRepository
	sessionIDGen        sessionid.Generator
}

func ProvideSessionmanager(ctx context.Context, repository repository.SessionRepository, sessionIDGen sessionid.Generator) *SessionHandler {
	return &SessionHandler{
		ctx:               ctx,
		sessionRepository: repository,
		sessionIDGen:      sessionIDGen,
	}
}

// route: /api/fatecore/start
func (mgr *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	mgr.startSession(w, r)
}

// route: /api/fatecore/join/{sessionid}
func (mgr *SessionHandler) JoinSession(w http.ResponseWriter, r *http.Request) {
	mgr.joinSession(w, r)
}

// route: /api/fatecore/leave/{sessionid}
func (mgr *SessionHandler) LeaveSession(w http.ResponseWriter, r *http.Request) {
	mgr.leaveSession(w, r)
}

func (mgr *SessionHandler) startSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusBadRequest)
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
		SessionRepository:  mgr.sessionRepository,
		SessionIDGenerator: mgr.sessionIDGen,
	}

	id, err := startsession.Start(mgr.ctx, startsessionPorts, startsession.Request{
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

func (mgr *SessionHandler) joinSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var characterID int
	if err := json.NewDecoder(r.Body).Decode(&characterID); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	joinSessionPorts := joinsession.Ports{
		SessionRepository:   mgr.sessionRepository,
		CharacterRepository: mgr.characterRepository,
	}

	err := joinsession.Join(mgr.ctx, joinSessionPorts, joinsession.Request{
		SessionID:   model.SessionID(r.PathValue("sessionid")),
		CharacterID: characterID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("error joining session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mgr *SessionHandler) leaveSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var characterID int
	if err := json.NewDecoder(r.Body).Decode(&characterID); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	leaveSessionPorts := leavesession.Ports{
		SessionRepository:   mgr.sessionRepository,
		CharacterRepository: mgr.characterRepository,
	}

	err := leavesession.Leave(mgr.ctx, leaveSessionPorts, leavesession.Request{
		SessionID:   model.SessionID(r.PathValue("sessionid")),
		CharacterID: characterID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("error leaving session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
