package twelve

import (
	"fmt"
	"strings"
)

func Verse(n int) string {
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	presents := []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

	var allPresents []string
	for i := n - 1; i >= 0; i-- {
		f := presents[i]
		if i == 0 && n >= 2 {
			f = fmt.Sprintf("and %s", presents[i])
		}

		allPresents = append(allPresents, f)
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[n - 1], strings.Join(allPresents, ", "))
}

func Song() string {
	var result []string
	for i := 0; i < 12; i++ {
		result = append(result, Verse(i+1))
	}
	return strings.Join(result, "\n")
}
