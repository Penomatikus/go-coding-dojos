package sessionid

import "github.com/Penomatikus/onionarch/internal/domain/model"

type Generator interface {
	GenerateSessionID() (model.SessionID, error)
}
