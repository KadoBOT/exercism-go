package pascal

func factorial(number int) uint64 {
	var result uint64 = 1
	if number < 0 {

	} else {
		for i := 1; i <= number; i++ {
			result *= uint64(i)
		}
	}
	return result
}

func Triangle(rows int) [][]int {
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		row := make([]int, i+1)
		rowFac := factorial(i)
		for j := 0; j <= i; j++ {
			colFac := factorial(j)
			row[j] = int(rowFac / (colFac * factorial(i-j)))
		}
		result[i] = row
	}
	return result
}
