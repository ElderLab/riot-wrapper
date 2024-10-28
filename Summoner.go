package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"encoding/json"
)

type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (s Summoner) String() string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// GetSummonerByPUUID is a function that returns a summoner by PUUID.
func (cli *RiotClient) GetSummonerByPUUID(puuid string) (*Summoner, []error) {
	result, errs := cli.lolClient.Get("/lol/summoner/v4/summoners/by-puuid/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var summoner Summoner
	err := json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, []error{err}
	}
	return &summoner, nil
}

// GetSummonerByAccountId is a function that returns a summoner by AccountId.
func (cli *RiotClient) GetSummonerByAccountId(accountId string) (*Summoner, []error) {
	result, errs := cli.lolClient.Get("/lol/summoner/v4/summoners/by-account/:accountId", []request.ParamURL{
		{
			Key:   "accountId",
			Value: accountId,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var summoner Summoner
	err := json.Unmarshal(result, &summoner)
	if err != nil {
		return nil, []error{err}
	}
	return &summoner, nil
}
