package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
)

// GetSummonerByPUUID is a function that returns a summoner by PUUID.
func (cli *RiotClient) GetSummonerByPUUID(puuid string) (*response.Summoner, error) {
	result, err := cli.lolClient.Get("/lol/summoner/v4/summoners/by-puuid/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return nil, err
	}
	var summoner response.Summoner
	err = json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, err
	}
	return &summoner, nil
}

// GetSummonerByAccountId is a function that returns a summoner by AccountId.
func (cli *RiotClient) GetSummonerByAccountId(accountId string) (*response.Summoner, error) {
	result, err := cli.lolClient.Get("/lol/summoner/v4/summoners/by-account/:accountId", []request.ParamURL{
		{
			Key:   "accountId",
			Value: accountId,
		},
	})
	if err != nil {
		return nil, err
	}
	var summoner response.Summoner
	err = json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, err
	}
	return &summoner, nil
}
