package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(white, black string) (bool, error) {
	if len(white) != 2 || len(black) != 2 {
		return false, errors.New("Invalid input")
	}

	if white == black {
		return false, errors.New("Invalid input. Both positions cannot be the same")
	}

	whiteX := int(white[0]) // X co-ordinate of white piece
	whiteY := int(white[1]) // Y co-ordinate of white piece
	blackX := int(black[0])
	blackY := int(black[1])

	if whiteX > 'h' || whiteX < 'a' || blackX > 'h' || blackX < 'a' ||
		whiteY < '1' || whiteY > '8' || blackY < '1' || blackY > '8' {
		return false, errors.New("Invalid input position")
	}

	if whiteX == blackX || whiteY == blackY || math.Abs(float64(whiteX-blackX)) == math.Abs(float64(whiteY-blackY)) {
		return true, nil
	}

	return false, nil

}
