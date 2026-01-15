package dictionary

import "github.com/dinhphu28/dictionary/database"

type LookupResult struct {
	ID         string           `json:"id"`
	Dictionary string           `json:"dictionary"`
	FullName   string           `json:"full_name"`
	Entries    []database.Entry `json:"entries"`
}

type MatchType int

const (
	Unknown MatchType = iota
	ExactMatch
	ApproximateMatch
)

type LookupResultWithSuggestion struct {
	LookupResults []LookupResult `json:"lookup_results"`
	MatchType     MatchType      `json:"match_type"`
	Suggestions   []string       `json:"suggestions"`
}
