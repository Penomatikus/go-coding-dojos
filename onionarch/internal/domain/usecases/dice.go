package usecases

import (
	"math/rand/v2"
)

// RollDice rolles 4x a fate dice and returns the total points
func RollDice() (result int) {
	faces := [...]int{-1, -1, 0, 0, 1, 1} // -, -, , , +, +
	for i := 0; i < 4; i++ {
		result += faces[rand.IntN(6)]
	}
	return
}
