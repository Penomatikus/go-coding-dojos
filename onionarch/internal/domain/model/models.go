package model

type (
	Session struct {
		SessionID string
		Game      Game
	}

	Game struct {
		Players []Player
	}

	Player struct {
		ID     int
		Type   int
		Name   string
		Points int
	}
)
