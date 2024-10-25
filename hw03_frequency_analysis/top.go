package hw03frequencyanalysis

import "math"

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	frequencyStat := frequencyCalc(s)

	wordSlice := getSortedSlice(frequencyStat)

	l := math.Min(10, float64(len(wordSlice)))
	return wordSlice[:int(l)]
}
