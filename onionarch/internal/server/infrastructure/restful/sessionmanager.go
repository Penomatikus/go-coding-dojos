package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/service"
)

type startRequest struct {
	Title   string `json:"title"`
	OwnerID int    `json:"owner_id"`
}

type SessionManager struct {
	ctx            context.Context
	sessionService service.Session
}

func ProvideSessionmanager(sessionService service.Session) *SessionManager {
	return &SessionManager{
		sessionService: sessionService,
	}
}

func (mgr *SessionManager) StartSession(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(http.MethodPost, r.Method) {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	var req startRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := mgr.sessionService.Start(context.Background(), req.Title, req.OwnerID); err != nil {
		http.Error(w, fmt.Sprintf("error creating new session: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
