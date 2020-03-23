package queenattack

import (
	"fmt"
)

func isValid(pos string) bool {
	invalidColumn := pos[0]-'a' < 0 || 'h'-pos[0] > 8
	invalidRow := pos[1]-'0' <= 0 || '8'-pos[1] > 8
	if len(pos) != 2 || invalidColumn || invalidRow {
		return false
	}

	return true
}
func CanQueenAttack(w string, b string) (bool, error) {
	if w == b {
		return false, fmt.Errorf("same square")
	}

	if !isValid(w) || !isValid(b) {
		return false, fmt.Errorf("invalid queen position")
	}

	col1, row1 := w[0], w[1]
	col2, row2 := b[0], b[1]

	switch {
	case col1 == col2 || row1 == row2:
		return true, nil
	case (col1-col2) == (row1-row2) || (col2-col1) == (row1-row2):
		return true, nil
	default:
		return false, nil
	}
}
