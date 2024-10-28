package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"ElderLab/riot-wrapper/models/response"
	"encoding/json"
)

// GetSummonerByPUUID is a function that returns a summoner by PUUID.
func (cli *RiotClient) GetSummonerByPUUID(puuid string) (*response.Summoner, []error) {
	result, errs := cli.lolClient.Get("/lol/summoner/v4/summoners/by-puuid/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var summoner response.Summoner
	err := json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, []error{err}
	}
	return &summoner, nil
}

// GetSummonerByAccountId is a function that returns a summoner by AccountId.
func (cli *RiotClient) GetSummonerByAccountId(accountId string) (*response.Summoner, []error) {
	result, errs := cli.lolClient.Get("/lol/summoner/v4/summoners/by-account/:accountId", []request.ParamURL{
		{
			Key:   "accountId",
			Value: accountId,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var summoner response.Summoner
	err := json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, []error{err}
	}
	return &summoner, nil
}
