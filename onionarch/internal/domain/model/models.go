package model

import "time"

type (
	SessionID string

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
		ID          int
		Name        string
		Description string
		SessionID   *SessionID
		PlayerID    int
		Points      int
	}

	Notification struct {
		CreatedAt time.Time
		SessionId SessionID
		FromId    int
		Body      string
	}
)
