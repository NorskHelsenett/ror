package fuzzysearch

import (
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func FuzzySearchandSortRanked(search string, list []string) []string {
	matches := fuzzy.RankFind(search, list)
	sort.Sort(matches)

	var sorted []string
	for _, match := range matches {
		sorted = append(sorted, match.Target)
	}
	return sorted
}
