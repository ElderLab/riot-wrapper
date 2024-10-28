package response

import "encoding/json"

// MatchesIds is a struct that represents a list of match IDs.
type MatchesIds []string

// String is a function that returns the matches IDs in json format.
func (m MatchesIds) String() string {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
