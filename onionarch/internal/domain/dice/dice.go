package dice

import "errors"

var ErrDiceRoll = errors.New("error while rolling dice")

type Dice[T any] interface {
	Roll() (T, error)
}
