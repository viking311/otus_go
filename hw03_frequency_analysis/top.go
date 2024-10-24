package hw03frequencyanalysis

func Top10(s string) []string {
	if len(s) == 0 {
		return nil
	}

	frequencyStat := frequencyCalc(s)

	wordSlice := getSortedSlice(frequencyStat)

	return wordSlice[:10]
}
