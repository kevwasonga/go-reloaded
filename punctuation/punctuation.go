package punctuation

import (
	"regexp"
	"strings"
)

func FormatPunctuation(text string) string {
	// Define regex patterns for punctuation rules
	punctuationPattern := `[.,!?;:]+`
	

	// Apply rules for  punctuation marks
	text = regexp.MustCompile(`\s*(` + punctuationPattern + `)\s*`).ReplaceAllString(text, "$1 ")
	

	// handling quotation marks
	text = regexp.MustCompile(`\s*'\s*([^']+?)\s*'\s*`).ReplaceAllString(text, " '$1' ")

	// Remove any remaining leading or trailing spaces
	text = strings.TrimSpace(text)

	return text
}
