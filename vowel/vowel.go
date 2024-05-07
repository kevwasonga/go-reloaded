package vowel

import (
	"strings"
)

// ApplyVowelModifications adjusts the usage of articles "a" and "an" according to the requirement
func ApplyVowelModifications(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && startsWithVowelOrH(words[i+1]) {
			words[i] = "an"
		} else if words[i] == "A" && startsWithVowelOrH(words[i+1]) {
			words[i] = "An"
		}
	}
	return strings.Join(words, " ")
}

func startsWithVowelOrH(word string) bool {
	firstChar := string(word[0])
	return strings.ContainsAny(firstChar, "aeiouAEIOU") || strings.EqualFold(firstChar, "h")
}
