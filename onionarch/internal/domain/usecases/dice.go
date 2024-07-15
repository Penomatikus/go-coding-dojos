package usecases

import (
	"math/rand/v2"

	"github.com/Penomatikus/onionarch/internal/domain/model"
)

// RollDice rolles 4x a fate dice and returns the total points
func RollDice() (result int) {
	for i := 0; i < 4; i++ {
		result += model.Dice()[rand.IntN(6)]
	}
	return
}
