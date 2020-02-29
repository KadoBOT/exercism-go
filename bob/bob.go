package bob

import "strings"

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	isYellingQuestion := isQuestion && isYell

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case isYellingQuestion:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYell:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
