package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	cols, rows [][]int
}


func getMatrixSlice(s [][]int) [][]int {
	slice := make([][]int, len(s))
	for i, nums := range s {
		for _, n := range nums {
			slice[i] = append(slice[i], n)
		}
	}
	return slice
}

func (m Matrix) Rows() [][]int {
	return getMatrixSlice(m.rows)
}

func (m Matrix) Cols() [][]int {
	return getMatrixSlice(m.cols)
}

func (m Matrix) Set(row, col, val int) bool {
	validRow := row >= 0 && row < len(m.rows)
	validCol := col >= 0 && col < len(m.cols)

	if validCol && validRow {
		m.rows[row][col] = val
		m.cols[col][row] = val
		return true
	}
	return false
}

func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")
	rows := make([][]int, len(lines))

	var cols [][]int
	for i, v := range lines {
		nums := strings.Split(v, " ")

		switch {
		case i ==0:
			cols = make([][]int, len(nums))
		case len(cols) != len(nums):
			return nil, fmt.Errorf("uneven rows")
		case len(nums) == 0:
			return nil, fmt.Errorf("empty row")
		}

		for j, num := range nums {
			n1, err := strconv.Atoi(num)

			if err != nil {
				return nil, fmt.Errorf("non-int value")
			}

			rows[i] = append(rows[i], n1)
			cols[j] = append(cols[j], n1)
		}
	}

	m := &Matrix{ cols, rows}
	return m, nil
}
