package utils

import (
	"errors"
	"net/http"
	"unicode/utf8"
)

// FetchSortParams. Fetches sort parameters from query string
// uses ?sort={operator}{field} - i.e. ?sort=-name - (name desc)
func FetchSortParams(r *http.Request, defaultField string, defaultOperator int) (*QuerySort, error) {
	var sort QuerySort

	// fetch QS
	q := r.URL.Query()
	if s, ok := q["sort"]; ok {
		if utf8.RuneCountInString(s[0]) < 2 {
			return nil, errors.New("Invalid sort criteria")
		}
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

	return &sort, nil
}
