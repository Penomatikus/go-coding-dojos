package db

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/service"
)

type dbStore struct {
	character map[int]*model.Character
	player    map[int]*model.Player
	session   map[string]*model.Session
}

type (
	characterRepository struct{ store *dbStore }
	playerRepository    struct{ store *dbStore }
	sessionRepository   struct{ store *dbStore }
)

var (
	_ service.CharacterRepository = &characterRepository{}
	_ service.PlayerRepository    = &playerRepository{}
	_ service.SessionRepository   = &sessionRepository{}
)

func ProvideSessionRepository(dbStore *dbStore) service.SessionRepository {
	return &sessionRepository{store: dbStore}
}

func (repo *sessionRepository) LoadSession(ctx context.Context, sessionID string) (*model.Session, error) {
	s, ok := repo.store.session[sessionID]
	if !ok {
		return nil, service.ErrNotFound
	}
	return s, nil
}

func (repo *sessionRepository) CreateSession(ctx context.Context, session *model.Session) error {
	if _, ok := repo.store.session[session.ID]; ok {
		return service.ErrAlreadyExists
	}

	repo.store.session[session.ID] = session
	return nil
}

func ProvidePlayerRepository(dbStore *dbStore) service.PlayerRepository {
	return &playerRepository{store: dbStore}
}

func (repo *playerRepository) SavePlayer(ctx context.Context, player *model.Player) error {
	if _, ok := repo.store.player[player.ID]; ok {
		return service.ErrAlreadyExists
	}

	repo.store.player[player.ID] = player
	return nil
}

func ProvideCharacterRepository(dbStore *dbStore) service.CharacterRepository {
	return &characterRepository{store: dbStore}
}

func (repo *characterRepository) SaveCharacter(ctx context.Context, character *model.Character) error {
	if _, ok := repo.store.character[character.ID]; ok {
		return service.ErrAlreadyExists
	}

	repo.store.character[character.ID] = character
	return nil
}
