package dictionary

import (
	"github.com/dinhphu28/dictionary/database"
	"github.com/dinhphu28/dictionary/pkg/ranking"
)

type ApproximateLookup struct {
	dictionaryLookup DictionaryLookup
}

func NewApproximateLookup(dictionaryLookup DictionaryLookup) *ApproximateLookup {
	return &ApproximateLookup{
		dictionaryLookup: dictionaryLookup,
	}
}

func (approximateLookup *ApproximateLookup) LookupWithSuggestion(q string) (LookupResultWithSuggestion, error) {
	results := approximateLookup.dictionaryLookup.LookupAllDictionariesAndSort(q)

	if len(results) > 0 {
		return LookupResultWithSuggestion{
			LookupResults: results,
			MatchType:     ExactMatch,
			Suggestions:   []string{},
		}, nil
	}

	// NOTE: Rule of thumb: use first 3–4 characters of query.
	prefix := q
	if len(prefix) > 4 {
		prefix = prefix[:4]
	}

	// NOTE: Prefer American English (id: oxford_american) database for suggestions.
	// Find the dictionary in the list.
	// TODO: Cover case where the preferred dictionary is not found.
	dictionaries := approximateLookup.dictionaryLookup.dictionaries
	var preferredDict *database.Dictionary
	for _, dict := range dictionaries {
		if dict.Manifest.ID == "oxford_american" {
			preferredDict = &dict
			break
		}
	}

	candidates, err := database.FindByHeadwordStartsWith(preferredDict.DB, prefix, 50)
	if err != nil {
		return LookupResultWithSuggestion{}, err
	}

	matches := ranking.RankByEditDistance(q, candidates)
	if len(matches) == 0 {
		return LookupResultWithSuggestion{}, nil
	}
	firstMatch := matches[0]

	secondaryResults := approximateLookup.dictionaryLookup.LookupAllDictionariesAndSort(
		firstMatch.Word,
	)

	suggestWord := []string{}
	for _, match := range matches {
		suggestWord = append(suggestWord, match.Word)
	}

	return LookupResultWithSuggestion{
		LookupResults: secondaryResults,
		MatchType:     ApproximateMatch,
		Suggestions:   suggestWord,
	}, nil
}
