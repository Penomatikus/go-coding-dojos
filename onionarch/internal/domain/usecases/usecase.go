package usecases

import "github.com/Penomatikus/onionarch/internal/domain/model"

type (
	SessionID func() model.SessionID
)
