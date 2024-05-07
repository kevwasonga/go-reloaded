package bin

import (
    "regexp"
    "strconv"
   
)

func ApplyBinModifications(text string) string {
    // Regex to identify (bin) and the binary number before it
    re := regexp.MustCompile(`(\b[01]+\b)\s*\((bin)\)`)

    // Replace instances of (bin) with the decimal version of the binary number before it
    modifiedText := re.ReplaceAllStringFunc(text, func(match string) string {
        parts := re.FindStringSubmatch(match)
        binaryWord := parts[1]
        decimalValue, err := strconv.ParseInt(binaryWord, 2, 64)
        if err != nil {
            // If parsing fails, return the original match
            return match
        }
        return strconv.FormatInt(decimalValue, 10)
    })

    return modifiedText
}
