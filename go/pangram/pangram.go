package pangram

import "strings"

const (
	testVersion       = 1
	letters           = "abcdefghijklmnopqrstuvwxyz"
	allLettersChecked = 26
)

func IsPangram(message string) bool {
	checkedLetters := make(map[rune]bool)
	lower := strings.ToLower(message)

	for _, runeChar := range lower {
		if strings.ContainsRune(letters, runeChar) {
			checkedLetters[runeChar] = true
		}
	}

	return len(checkedLetters) == allLettersChecked
}
