package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(message string) string {
	words := regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+").FindAllString(message, -1)
	acronym := make([]string, len(words))

	for _, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		acronym = append(acronym, firstLetter)
	}

	return strings.Join(acronym, "")
}
