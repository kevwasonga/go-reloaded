package up

import (
    "strconv"
    "strings"
)

// ApplyUpModifications capitalizes the specified number of words before "(up)"
func ApplyUpModifications(words string) string {
    data := strings.Split(words, " ")

    for i := 0; i < len(data); i++ {
        if strings.Contains(data[i], "up,") {
            // Extract the number after "up, "
            numStr := strings.TrimSuffix(strings.TrimPrefix(data[i+1], " "), ")")

            toInt, err := strconv.Atoi(numStr)
            if err != nil {
                // If there's an error converting to integer, skip this modification
                continue
            }

            // Capitalize the specified number of words before "(up)"
            for j := i - toInt; j < i; j++ {
                if j >= 0 && j < len(data) {
                    data[j] = strings.ToUpper(data[j])
                }
            }

            // Remove "(up)" and the number from the slice
            data = append(data[:i], data[i+2:]...)
            i-- // Adjust the index after removing elements
        } else if data[i] == "(up)" && i > 0 {
            // Convert to uppercase the word before "(up)"
            data[i-1] = strings.ToUpper(data[i-1])

            // Remove "(up)" from the slice
            data = append(data[:i], data[i+1:]...)
            i-- // Adjust the index after removing elements
        }
    }
    return strings.Join(data, " ")
}
