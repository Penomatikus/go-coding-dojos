package repositorytest

import (
	"context"
	"fmt"

	"github.com/Penomatikus/onionarch/internal/domain/model"
	"github.com/Penomatikus/onionarch/internal/domain/repository"
)

type dbStore struct {
	Character map[int]*model.Character
	Session   map[model.SessionID]*model.Session
}

func NewDBStore() dbStore {
	return dbStore{
		Character: make(map[int]*model.Character),
		Session:   make(map[model.SessionID]*model.Session),
	}
}

const _CHARACTER = 0

func (db *dbStore) autoIncrement(tableType int) int {
	switch tableType {
	case _CHARACTER:
		return len(db.Character) + 1
	default:
		panic(fmt.Sprintf("unkown table type %d", tableType))
	}
}

type (
	characterRepository struct{ store *dbStore }
	sessionRepository   struct{ store *dbStore }
)

var (
	_ repository.CharacterRepository = &characterRepository{}
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

func ProvideCharacterRepository(dbStore *dbStore) repository.CharacterRepository {
	return &characterRepository{store: dbStore}
}

func (repo *characterRepository) Create(ctx context.Context, Character *model.Character) (int, error) {
	Character.ID = repo.store.autoIncrement(_CHARACTER)
	if _, ok := repo.store.Character[Character.ID]; ok {
		return -1, repository.ErrAlreadyExists
	}

	repo.store.Character[Character.ID] = Character
	return Character.ID, nil
}

func (repo *characterRepository) FindByID(ctx context.Context, characterID int) (*model.Character, error) {
	c, ok := repo.store.Character[characterID]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return c, nil
}

func (repo *characterRepository) FindBySession(ctx context.Context, sessionID model.SessionID) ([]model.Character, error) {
	characters := make([]model.Character, 0)
	for _, char := range repo.store.Character {
		if *char.SessionID == sessionID {
			characters = append(characters, *char)
		}
	}
	return characters, nil
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
