package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"ElderLab/riot-wrapper/models/response"
	"encoding/json"
)

// GetMatchesIds is a function that returns a list of match IDs by PUUID.
func (cli *RiotClient) GetMatchesIds(puuid string) (*response.MatchesIds, []error) {
	result, errs := cli.riotClient.Get("/lol/match/v5/matches/by-puuid/:puuid/ids", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var matchesIds response.MatchesIds
	err := json.Unmarshal(result, &matchesIds)
	if err != nil {
		return nil, []error{err}
	}
	return &matchesIds, nil
}

// GetMatchById is a function that returns a match details by match ID.
func (cli *RiotClient) GetMatchById(matchId string) (string, []error) {
	result, err := cli.riotClient.Get("/lol/match/v5/matches/:matchId", []request.ParamURL{
		{
			Key:   "matchId",
			Value: matchId,
		},
	})
	if err != nil {
		return "", err
	}
	return string(result), nil
}
