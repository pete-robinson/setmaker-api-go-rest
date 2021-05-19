package utils

import "net/http"

// FetchSortParams. Fetches sort parameters from query string
// uses ?sort={operator}{field} - i.e. ?sort=-name - (name desc)
func FetchSortParams(r *http.Request, defaultField string, defaultOperator int) *QuerySort {
	var sort QuerySort

	// fetch QS
	q := r.URL.Query()
	if s, ok := q["sort"]; ok {
		// split to runes
		sRune := []rune(s[0])
		operator := string(sRune[:1])
		sort.Field = string(sRune[1:])
		sort.Operator = 1

		// overwrite operator to -1 if operator is -
		if operator == "-" {
			sort.Operator = -1
		}
	} else {
		// fallback to default fields
		sort.Field = defaultField
		sort.Operator = defaultOperator
	}

	return &sort
}
