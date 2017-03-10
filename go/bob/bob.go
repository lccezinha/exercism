package bob

import "strings"

const testVersion = 2

func Hey(message string) string {
	message = strings.TrimSpace(message)

	if len(message) == 0 {
		return "Fine. Be that way!"
	}

	if isShouting(message) {
		return "Whoa, chill out!"
	}

	if isQuestion(message) {
		return "Sure."
	}

	return "Whatever."
}

func isQuestion(message string) bool {
	return strings.HasSuffix(message, "?")
}

func isShouting(message string) bool {
	return message == strings.ToUpper(message) && message != strings.ToLower(message)
}
