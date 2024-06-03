package model

import "time"

type (
	SessionID   string
	CharacterID int

	Session struct {
		ID        SessionID
		Owner     int
		CreatedAt time.Time
		Title     string
	}

	Player struct {
		ID        int
		CreatedAt time.Time
		Name      string
	}

	Character struct {
		ID        CharacterID
		SessionID *SessionID
		PlayerID  int
		Points    int
	}
)
