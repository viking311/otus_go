package hw03frequencyanalysis

import "sort"

func getSortedSlice(frequency map[string]int) []string {
	pl := make(PairList, len(frequency))
	i := 0
	for k, v := range frequency {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pl))

	wordSlice := make([]string, 0, len(pl))
	for _, v := range pl {
		wordSlice = append(wordSlice, v.Key)
	}

	return wordSlice
}
