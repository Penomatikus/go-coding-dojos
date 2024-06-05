package sessionid

import (
	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/sessionid"
	"github.com/gofrs/uuid"
)

type SessionID struct{}

func (sid *SessionID) GenerateSessionID() (model.SessionID, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return model.SessionID(id.String()), nil
}

func ProvideSessionIDGen() sessionid.Generator {
	return &SessionID{}
}
