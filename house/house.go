package house

import (
	"fmt"
	"strings"
)

var prefixes = []string{
	"",
	"belonged to",
	"kept",
	"woke",
	"married",
	"kissed",
	"milked",
	"tossed",
	"worried",
	"killed",
	"ate",
	"lay in",
}

var sentences = []string{
	"the horse and the hound and the horn",
	"the farmer sowing his corn",
	"the rooster that crowed in the morn",
	"the priest all shaven and shorn",
	"the man all tattered and torn",
	"the maiden all forlorn",
	"the cow with the crumpled horn",
	"the dog",
	"the cat",
	"the rat",
	"the malt",
	"the house that Jack built.",
}

func Verse(v int) string {
	var result string
	for i := v; i > 0; i-- {
		sentence := sentences[len(sentences)-i]
		prefix := prefixes[len(prefixes)-i]
		if i == v {
			result += fmt.Sprintf("This is %s", sentence)
		} else {
			result += fmt.Sprintf("that %s %s", prefix, sentence)
		}

		if i != 1 {
			result += "\n"
		}
	}

	return result
}

func Song() string {
	var verses []string
	n := len(sentences)
	for i := 1; i <= n; i++ {
		v := Verse(i)
		verses = append(verses, v)
	}
	return strings.Join(verses, "\n\n")
}
