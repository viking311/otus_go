package hw03frequencyanalysis

import "strings"

func frequencyCalc(text string) map[string]int {
	rawWords := strings.Fields(text)

	frequency := make(map[string]int)
	for _, word := range rawWords {
		normWord := normalizeWord(word)
		if len(normWord) > 0 {
			frequency[normWord]++
		}
	}

	return frequency
}

func normalizeWord(word string) string {
	t := strings.Trim(word, "\t\n .,?!-\":;")
	return strings.ToLower(t)
}
