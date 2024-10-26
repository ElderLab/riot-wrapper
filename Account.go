package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"encoding/json"
)

// Account is a struct that represents the account riot.
type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

// String is a function that returns the account in json format.
func (a Account) String() string {
	//print the account to json format formated with tabulation
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// GetAccountByRiotId is a function that returns the account ID for a given gameName and tagLine.
func (cli *RiotClient) GetAccountByRiotId(gameName, tagLine string) (*Account, []error) {
	result, errs := cli.riotClient.Get("/riot/account/v1/accounts/by-riot-id/:gameName/:tagLine", []request.ParamURL{
		{
			Key:   "gameName",
			Value: gameName,
		},
		{
			Key:   "tagLine",
			Value: tagLine,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var account Account
	err := json.Unmarshal(result, &account)
	if err != nil {
		return nil, []error{err}
	}
	return &account, nil
}
