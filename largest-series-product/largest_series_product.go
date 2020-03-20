package lsproduct

import (
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, series int) (int, error) {
	if series > len(digits) || series < 0 {
		return 0, fmt.Errorf("invalid argument provided")
	}

	if digits == "" || series == 0 {
		return 1, nil
	}

	var result int
	for i := 0; i < len(digits)-series+1; i++ {
		digitSlice := digits[i : i+series]
		temp, err := strconv.Atoi(string(digitSlice[0]))

		if err != nil {
			return 0, err
		}
		for _, v := range digitSlice[1:] {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				return 0, err
			}
			temp *= n
		}
		if temp > result {
			result = temp
		}
	}
	return result, nil
}
