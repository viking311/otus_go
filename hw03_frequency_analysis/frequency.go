package hw03frequencyanalysis

import "strings"

func frequencyCalc(text string) map[string]int {
	frequency := make(map[string]int)
	word := ""
	for _, r := range text {
		if r == ' ' || r == '\t' || r == '\n' {
			normWord := normalizeWord(word)
			if len(normWord) > 0 {
				frequency[normWord]++
			}
			word = ""
			continue
		}
		word += string(r)
	}

	if len(word) > 0 {
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
