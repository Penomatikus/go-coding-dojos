package repositorytest

import (
	"context"
	"fmt"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type dbStore struct {
	Character map[int]*model.Character
	Player    map[int]*model.Player
	Session   map[model.SessionID]*model.Session
}

func NewDBStore() dbStore {
	return dbStore{
		Character: make(map[int]*model.Character),
		Player:    make(map[int]*model.Player),
		Session:   make(map[model.SessionID]*model.Session),
	}
}

const (
	_CHARACTER = iota
	_PLAYER
)

func (db *dbStore) autoIncrement(tableType int) int {
	switch tableType {
	case _CHARACTER:
		return len(db.Character) + 1
	case _PLAYER:
		return len(db.Player) + 1
	default:
		panic(fmt.Sprintf("unkown table type %d", tableType))
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

func (repo *playerRepository) Create(ctx context.Context, player *model.Player) error {
	player.ID = repo.store.autoIncrement(_PLAYER)
	if _, ok := repo.store.Player[player.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.Player[player.ID] = player
	return nil
}

func (repo *playerRepository) Update(ctx context.Context, player *model.Player) error {
	p, ok := repo.store.Player[player.ID]
	if !ok {
		return repository.ErrNotFound
	}

	p.Name = player.Name
	repo.store.Player[player.ID] = p
	return nil
}

func (repo *playerRepository) FindByID(ctx context.Context, ID int) (*model.Player, error) {
	p, ok := repo.store.Player[ID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return p, nil
}

func ProvideCharacterRepository(dbStore *dbStore) repository.CharacterRepository {
	return &characterRepository{store: dbStore}
}

func (repo *characterRepository) Create(ctx context.Context, Character *model.Character) error {
	Character.ID = repo.store.autoIncrement(_CHARACTER)
	if _, ok := repo.store.Character[Character.ID]; ok {
		return repository.ErrAlreadyExists
	}

	repo.store.Character[Character.ID] = Character
	return nil
}

func (repo *characterRepository) FindByID(ctx context.Context, characterID int) (*model.Character, error) {
	c, ok := repo.store.Character[characterID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return c, nil
}

func (repo *characterRepository) Update(ctx context.Context, character *model.Character) error {
	c, ok := repo.store.Character[character.ID]
	if !ok {
		return repository.ErrNotFound
	}

	c.Points = character.Points
	c.SessionID = character.SessionID
	repo.store.Character[character.ID] = c
	return nil
}
