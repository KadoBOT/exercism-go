package raindrops

import "strconv"

func Convert(num int) (result string)  {
	if num % 3 == 0 {
		result += "Pling"
	}
	if num % 5 == 0 {
		result += "Plang"
	}
	if num % 7 == 0 {
		result += "Plong"
	}
	if result == ""{
		result = strconv.Itoa(num)
	}
	return
}