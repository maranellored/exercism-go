package grains

import (
	"errors"
)

func Square(input int) (output uint64, err error) {
	if input > 64 || input < 1 {
		return 0, errors.New("invalid input")
	}
	return 1 << uint(input-1), nil
}

func Total() uint64 {
	// This is a power series of the form
	// x^0 + x^1 + x^2 + x^3 .... + x^63
	// Here x is 2.
	return 1<<64 - 1
}
