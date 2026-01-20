package ranking

import (
	"sort"
	"strings"
)

type Match struct {
	Word string
	Dist int
}

func RankByEditDistanceWithMaxDist(query string, candidates []string, maxDist int) []Match {
	matches := RankByEditDistance(query, candidates)

	matchesByMaxDist := []Match{}

	for _, match := range matches {
		if match.Dist <= maxDist {
			matchesByMaxDist = append(matchesByMaxDist, match)
		}
	}

	return matchesByMaxDist
}

func RankByEditDistance(query string, candidates []string) []Match {
	q := strings.ToLower(query)

	var list []Match
	for _, c := range candidates {
		d := Levenshtein(q, strings.ToLower(c))
		list = append(list, Match{Word: c, Dist: d})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Dist == list[j].Dist {
			return list[i].Word < list[j].Word
		}
		return list[i].Dist < list[j].Dist
	})

	return list
}
