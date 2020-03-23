package pascal

func Triangle(rows int) [][]int {
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
