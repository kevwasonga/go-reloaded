package low

import (
    "strconv"
    "strings"
)

// ApplyLowModifications converts the specified number of words before "(low)" to lowercase
func ApplyLowModifications(words string) string {
    data := strings.Split(words, " ")

    for i := 0; i < len(data); i++ {
        if strings.Contains(data[i], "low,") {
            // Extract the number after "low, "
            numStr := strings.TrimSuffix(data[i+1], ")")

            toInt, err := strconv.Atoi(numStr)
            if err != nil {
                // If there's an error converting to integer, skip this modification
                continue
            }

            // Convert the specified number of words before "(low)" to lowercase
            for j := i - toInt; j < i; j++ {
                if j >= 0 && j < len(data) {
                    data[j] = strings.ToLower(data[j])
                }
            }

            // Remove "(low)" and the number from the slice
            data = append(data[:i], data[i+2:]...)
            i-- // Adjust the index after removing elements
        } else if data[i] == "(low)" && i > 0 {
            // Convert to lowercase the word before "(low)"
            data[i-1] = strings.ToLower(data[i-1])

            // Remove "(low)" from the slice
            data = append(data[:i], data[i+1:]...)
            i-- // Adjust the index after removing elements
        }
    }
    return strings.Join(data, " ")
}
