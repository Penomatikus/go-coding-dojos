package model

import "time"

type (
	Session struct {
		ID         string
		Characters []Character
		CreatedAt  time.Time
		Title      string
	}

	Player struct {
		ID        int
		CreatedAt time.Time
		Name      string
	}

	Character struct {
		ID       int
		PlayerID int
		Points   int
	}
)
