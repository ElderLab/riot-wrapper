package response

import "encoding/json"

// Account is a struct that represents the account riot.
type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

// String is a function that returns the account in json format.
func (a Account) String() string {
	//print the account to json format formatted with tabulation
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
