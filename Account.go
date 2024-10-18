package riot_wrapper

import "ElderLab/riot-wrapper/internal/request"

// GetAccountByRiotId is a function that returns the account ID for a given gameName and tagLine.
func (cli *RiotClient) GetAccountByRiotId(gameName, tagLine string) (string, []error) {

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
		return "", err
	}
	return string(result), nil
}
