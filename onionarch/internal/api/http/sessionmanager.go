package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
	"github.com/Penomatikus/onionarch/internal/domain/usecases/startsession"
)

type startRequest struct {
	Title   string `json:"title"`
	OwnerID int    `json:"owner_id"`
}

type SessionManager struct {
	ctx        context.Context
	repository repository.SessionRepository
}

func ProvideSessionmanager(ctx context.Context, repository repository.SessionRepository) *SessionManager {
	return &SessionManager{
		ctx: ctx,
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

	id, err := startsession.Start(mgr.ctx, startsession.Ports{SessionRepo: mgr.repository}, startsession.Request{
		ID:      func() model.SessionID { return "TODO " },
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
