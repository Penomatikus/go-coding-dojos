package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/startsession"
)

type startRequest struct {
	Title   string `json:"title"`
	OwnerID int    `json:"owner_id"`
}

type SessionHandler struct {
	ctx          context.Context
	repository   repository.SessionRepository
	sessionIDGen sessionid.Generator
}

func ProvideSessionmanager(ctx context.Context, repository repository.SessionRepository, sessionIDGen sessionid.Generator) *SessionHandler {
	return &SessionHandler{
		ctx:          ctx,
		repository:   repository,
		sessionIDGen: sessionIDGen,
	}
}

func (mgr *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var req startRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	startsessionPorts := startsession.Ports{
		SessionRepository:  mgr.repository,
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

func (mgr *SessionHandler) JoinSession(w *http.ResponseWriter, r *http.Request) {

}
