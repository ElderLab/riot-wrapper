package riot_wrapper

import (
	"ElderLab/riot-wrapper/internal/request"
	"encoding/json"
)

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

// GetMatchesIds is a function that returns a list of match IDs by PUUID.
func (cli *RiotClient) GetMatchesIds(puuid string) (*MatchesIds, []error) {
	result, errs := cli.riotClient.Get("/lol/match/v5/matches/by-puuid/:puuid/ids", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if errs != nil {
		return nil, errs
	}
	var matchesIds MatchesIds
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
