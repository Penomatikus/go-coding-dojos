package service

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
		CreateSession(ctx context.Context, session *model.Session) error
		LoadSession(ctx context.Context, sessionID string) (*model.Session, error)
	}

	PlayerRepository interface {
		SavePlayer(ctx context.Context, player *model.Player) error
	}

	CharacterRepository interface {
		SaveCharacter(ctx context.Context, character *model.Character) error
	}
)
