package hex

import (
	"fmt"
	"regexp"
	"strconv"
)

// ApplyHexModifications replaces the word before (hex) with its decimal equivalent
func ApplyHexModifications(text string) string {
	// Regular expression pattern to identify (hex) and the hexadecimal number before it
	re := regexp.MustCompile(`(\b[\da-fA-F]+\b)\s*\((hex)\)`)

	// Replace instances of (hex) with the decimal version of the hexadecimal number before it
	modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := re.FindStringSubmatch(match)
		hexWord := parts[1]
		decimalValue, err := strconv.ParseInt(hexWord, 16, 64)
		if err != nil {
			// If parsing fails, return the original match
			return match
		}
		return fmt.Sprintf("%d", decimalValue)
	})

	return modifiedText
}

