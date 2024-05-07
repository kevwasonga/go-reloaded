package caps

import (
	
	"strconv"
	"strings"

)

func ApplyCapModifications(words string) string {
	data := strings.Split(words, " ")

	for i := 0; i < len(data); i++ {
		
			if strings.Contains(data[i], "cap,") {
				// Extract the number after "cap, "
				numStr := strings.TrimSuffix(strings.TrimPrefix(data[i+1], " "), ")")

				toInt, err := strconv.Atoi(numStr)
				if err != nil {
					// If there's an error converting to integer, skip this modification
					continue
				}

				// Capitalize the specified number of words before "(cap)"
				for j := i - toInt; j < i; j++ {
					if j >= 0 && j < len(data) {
						data[j] = strings.Title(data[j])
					}
				}
			

			// Remove "(cap)" and the number from the slice
			data = append(data[:i], data[i+2:]...)
			i-- // Adjust the index after removing elements
		} else if data[i] == "(cap)" && i > 0 {
			// Capitalize the word before "(cap)"
			data[i-1] = strings.Title(data[i-1])

			// Remove "(cap)" from the slice
			data = append(data[:i], data[i+1:]...)
			i-- // Adjust the index after removing elements
		}
	}
	return strings.Join(data, " ")
}


