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
		Update(ctx context.Context, session *model.Session) error
		Create(ctx context.Context, session *model.Session) error
		FindByID(ctx context.Context, sessionID model.SessionID) (*model.Session, error)
	}

	PlayerRepository interface {
		Save(ctx context.Context, player *model.Player) error
	}

	CharacterRepository interface {
		Update(ctx context.Context, character *model.Character) error
		Create(ctx context.Context, character *model.Character) error
		FindByID(ctx context.Context, characterID model.CharacterID) (*model.Character, error)
	}
)
