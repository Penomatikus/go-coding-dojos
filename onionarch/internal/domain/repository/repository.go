package repository

import (
	"context"
	"errors"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type (
	SessionRepository interface {
		Create(ctx context.Context, session *model.Session) error
		Update(ctx context.Context, session *model.Session) error
		FindByID(ctx context.Context, sessionID model.SessionID) (*model.Session, error)
	}

	CharacterRepository interface {
		Create(ctx context.Context, character *model.Character) (int, error)
		Update(ctx context.Context, character *model.Character) error
		FindByID(ctx context.Context, ID int) (*model.Character, error)
		FindBySession(ctx context.Context, sessionID model.SessionID) ([]model.Character, error)
	}
)
