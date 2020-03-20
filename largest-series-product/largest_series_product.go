package lsproduct

import (
	"fmt"
)

func LargestSeriesProduct(digits string, series int) (int, error) {
	if series > len(digits) || series < 0 {
		return 0, fmt.Errorf("invalid argument provided")
	}

	if digits == "" || series == 0 {
		return 1, nil
	}

	var result int
	for i := 0; i <= len(digits)-series; i++ {
		digitSlice := digits[i : i+series]
		if digitSlice[0] < '0' || digitSlice[0] > '9' {
			return 0, fmt.Errorf("digits contain non digit")
		}
		temp := int(digitSlice[0] - '0')

		for _, v := range digitSlice[1:] {
			if v < '0' || v > '9' {
				return 0, fmt.Errorf("digits contain non digit")
			}
			temp *= int(v - '0')
		}
		if temp > result {
			result = temp
		}
	}
	return result, nil
}
