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

	PlayerRepository interface {
		Create(ctx context.Context, player *model.Player) error
		Update(ctx context.Context, player *model.Player) error
		FindByID(ctx context.Context, ID int) (*model.Player, error)
	}

	CharacterRepository interface {
		Create(ctx context.Context, character *model.Character) error
		Update(ctx context.Context, character *model.Character) error
		FindByID(ctx context.Context, characterID model.CharacterID) (*model.Character, error)
	}
)
