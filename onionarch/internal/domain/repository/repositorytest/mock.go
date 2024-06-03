package repositorytest

import (
	"context"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type dbStore struct {
	Character map[model.CharacterID]*model.Character
	Player    map[int]*model.Player
	Session   map[model.SessionID]*model.Session
}

func NewDBStore() dbStore {
	return dbStore{
		Character: make(map[model.CharacterID]*model.Character),
		Player:    make(map[int]*model.Player),
		Session:   make(map[model.SessionID]*model.Session),
	}
}

type (
	CharacterRepository struct{ store *dbStore }
	playerRepository    struct{ store *dbStore }
	sessionRepository   struct{ store *dbStore }
)

var (
	_ repository.CharacterRepository = &CharacterRepository{}
	_ repository.PlayerRepository    = &playerRepository{}
	_ repository.SessionRepository   = &sessionRepository{}
)

func ProvideSessionRepository(dbStore *dbStore) repository.SessionRepository {
	return &sessionRepository{store: dbStore}
}

func (repo *sessionRepository) FindByID(ctx context.Context, sessionID model.SessionID) (*model.Session, error) {
	s, ok := repo.store.Session[sessionID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return s, nil
}

func (repo *sessionRepository) Create(ctx context.Context, session *model.Session) error {
	if _, ok := repo.store.Session[session.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.Session[session.ID] = session
	return nil
}

func (repo *sessionRepository) Update(ctx context.Context, session *model.Session) error {
	s, ok := repo.store.Session[session.ID]
	if !ok {
		return repository.ErrNotFound
	}

	repo.store.Session[session.ID] = &model.Session{
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
	if _, ok := repo.store.Player[player.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.Player[player.ID] = player
	return nil
}

func ProvideCharacterRepository(dbStore *dbStore) repository.CharacterRepository {
	return &CharacterRepository{store: dbStore}
}

func (repo *CharacterRepository) Create(ctx context.Context, Character *model.Character) error {
	if _, ok := repo.store.Character[Character.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.Character[Character.ID] = Character
	return nil
}

func (repo *CharacterRepository) FindByID(ctx context.Context, characterID model.CharacterID) (*model.Character, error) {
	c, ok := repo.store.Character[characterID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return c, nil
}

func (repo *CharacterRepository) Update(ctx context.Context, character *model.Character) error {
	_, ok := repo.store.Character[character.ID]
	if !ok {
		return repository.ErrNotFound
	}
	return nil
}
