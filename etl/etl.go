package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	if input == nil {
		return nil
	}

	var output map[string]int
	output = make(map[string]int)
	for value, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = value
		}
	}

	return output
}
