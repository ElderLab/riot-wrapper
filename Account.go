package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
)

// GetAccountByRiotId is a function that returns the account ID for a given gameName and tagLine.
func (cli *RiotClient) GetAccountByRiotId(gameName, tagLine string) (*response.Account, error) {
	result, err := cli.riotClient.Get("/riot/account/v1/accounts/by-riot-id/:gameName/:tagLine", []request.ParamURL{
		{
			Key:   "gameName",
			Value: gameName,
		},
		{
			Key:   "tagLine",
			Value: tagLine,
		},
	})
	if err != nil {
		return nil, err
	}
	var account response.Account
	err = json.Unmarshal(result, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
