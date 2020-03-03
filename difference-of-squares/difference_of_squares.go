package diffsquares

import "math"

func getResult(n int) (int, int) {
	var squareOfSum, sumOfSquares int
	for i:=1; i<=n; i++{
		squareOfSum += i
		sumOfSquares += i * i
	}
	return int(math.Pow(float64(squareOfSum), 2)), sumOfSquares
}

func SquareOfSum(n int) int {
	squareOfSum, _ := getResult(n)
	return squareOfSum
}

func SumOfSquares(n int) int {
	_, sumOfSquares := getResult(n)
	return sumOfSquares
}

func Difference(n int) int {
	squareOfSum, sumOfSquares := getResult(n)
	return squareOfSum - sumOfSquares
}
