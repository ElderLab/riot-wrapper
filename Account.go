package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"ElderLab/riot-wrapper/models/response"
	"encoding/json"
)

// GetAccountByRiotId is a function that returns the account ID for a given gameName and tagLine.
func (cli *RiotClient) GetAccountByRiotId(gameName, tagLine string) (*response.Account, []error) {
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
	var account response.Account
	err := json.Unmarshal(result, &account)
	if err != nil {
		return nil, []error{err}
	}
	return &account, nil
}
