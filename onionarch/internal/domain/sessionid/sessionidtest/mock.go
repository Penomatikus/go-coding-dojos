package sessionidtest

import (
	"time"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
)

type SessionIDGen struct{}

func (sig *SessionIDGen) GenerateSessionID() (model.SessionID, error) {
	return model.SessionID(time.Now().Format(time.RFC3339)), nil

}

func ProvideSessionIDGen() sessionid.Generator {
	return &SessionIDGen{}
}
