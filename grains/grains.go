package grains

import (
	"errors"
)

func Square(s int) (uint64, error) {
	if s > 64 || s < 1 {
		return 0, errors.New("invalid number of squares")
	}
	return 1 << uint(s - 1), nil
}

func Total() uint64 {
	return ^uint64(0)
}
