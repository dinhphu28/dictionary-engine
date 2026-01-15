package api

import (
	"net/http"
	"strings"

	"dinhphu28.com/dictionary"
)

type LookupHandlerV2 struct{}

func NewLookupHandlerV2() *LookupHandlerV2 {
	return &LookupHandlerV2{}
}

func (lookupHandler *LookupHandlerV2) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q == "" {
		http.Error(w, "missing q parameter", http.StatusBadRequest)
		return
	}

	result, err := dictionary.Lookup(q)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(result.LookupResults) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	writeJSONResponse(w, result)
}
