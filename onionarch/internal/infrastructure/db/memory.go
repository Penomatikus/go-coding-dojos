package db

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type dbStore struct {
	character map[model.CharacterID]*model.Character
	player    map[int]*model.Player
	session   map[model.SessionID]*model.Session
}

func NewDBStore() dbStore {
	return dbStore{
		character: make(map[model.CharacterID]*model.Character),
		player:    make(map[int]*model.Player),
		session:   make(map[model.SessionID]*model.Session),
	}
}

type (
	characterRepository struct{ store *dbStore }
	playerRepository    struct{ store *dbStore }
	sessionRepository   struct{ store *dbStore }
)

var (
	_ repository.CharacterRepository = &characterRepository{}
	_ repository.PlayerRepository    = &playerRepository{}
	_ repository.SessionRepository   = &sessionRepository{}
)

func ProvideSessionRepository(dbStore *dbStore) repository.SessionRepository {
	return &sessionRepository{store: dbStore}
}

func (repo *sessionRepository) FindByID(ctx context.Context, sessionID model.SessionID) (*model.Session, error) {
	s, ok := repo.store.session[sessionID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return s, nil
}

func (repo *sessionRepository) Create(ctx context.Context, session *model.Session) error {
	if _, ok := repo.store.session[session.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.session[session.ID] = session
	return nil
}

func (repo *sessionRepository) Update(ctx context.Context, session *model.Session) error {
	s, ok := repo.store.session[session.ID]
	if !ok {
		return repository.ErrNotFound
	}

	repo.store.session[session.ID] = &model.Session{
		ID:        session.ID,
		CreatedAt: s.CreatedAt,
		Owner:     session.Owner,
		Title:     session.Title,
	}

	return nil
}

func ProvidePlayerRepository(dbStore *dbStore) repository.PlayerRepository {
	return &playerRepository{store: dbStore}
}

func (repo *playerRepository) Save(ctx context.Context, player *model.Player) error {
	if _, ok := repo.store.player[player.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.player[player.ID] = player
	return nil
}

func ProvideCharacterRepository(dbStore *dbStore) repository.CharacterRepository {
	return &characterRepository{store: dbStore}
}

func (repo *characterRepository) Create(ctx context.Context, character *model.Character) error {
	if _, ok := repo.store.character[character.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.character[character.ID] = character
	return nil
}

func (repo *characterRepository) FindByID(ctx context.Context, characterID model.CharacterID) (*model.Character, error) {
	c, ok := repo.store.character[characterID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return c, nil
}

func (repo *characterRepository) Update(ctx context.Context, character *model.Character) error {
	_, ok := repo.store.character[character.ID]
	if !ok {
		return repository.ErrNotFound
	}
	return nil
}