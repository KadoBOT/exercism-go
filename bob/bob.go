package bob

import "strings"

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	isYell := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	isYellingQuestion := isQuestion && isYell

	if isYellingQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}
	if isYell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
