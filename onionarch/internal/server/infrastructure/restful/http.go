package restful

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penomatikus/onionarch/internal/domain/service"
)

type SessionHandler struct {
	sessionService service.Session
}

func NewSessionHandler(sessionService service.Session) SessionHandler {
	return SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) StartSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := h.sessionService.New(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, sessionID)
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/start", nil)
	mux.HandleFunc("POST /api/join/{sessionID}", nil)
	return mux
}

// cheap convinient method
func methodAllowed(want, got string) bool {
	return want != got
}
