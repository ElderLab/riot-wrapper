package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/opts"
	"github.com/ElderLab/riot-wrapper/models/response"
)

// GetMatchesIds is a function that returns a list of match IDs by PUUID.
func (cli *RiotClient) GetMatchesIds(puuid string) (*response.MatchesIds, error) {
	result, err := cli.riotClient.Get("/lol/match/v5/matches/by-puuid/:puuid/ids", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return nil, err
	}
	var matchesIds response.MatchesIds
	err = json.Unmarshal(result, &matchesIds)
	if err != nil {
		return nil, err
	}
	return &matchesIds, nil
}

// GetMatchesIdsWithOpts is a function that returns a list of match IDs by PUUID with options.
func (cli *RiotClient) GetMatchesIdsWithOpts(puuid string, opts opts.MatchesIdsOpts) (*response.MatchesIds, error) {
	params := request.StructToListOfParamURL(opts)
	params = append(params, request.ParamURL{
		Key:   "puuid",
		Value: puuid,
	})
	result, err := cli.riotClient.Get("/lol/match/v5/matches/by-puuid/:puuid/ids", params)
	if err != nil {
		return nil, err
	}
	var matchesIds response.MatchesIds
	err = json.Unmarshal(result, &matchesIds)
	if err != nil {
		return nil, err
	}
	return &matchesIds, nil
}

// GetMatchById is a function that returns a match details by match ID.
func (cli *RiotClient) GetMatchById(matchId string) (*response.Match, error) {
	result, err := cli.riotClient.Get("/lol/match/v5/matches/:matchId", []request.ParamURL{
		{
			Key:   "matchId",
			Value: matchId,
		},
	})
	if err != nil {
		return nil, err
	}
	var match response.Match
	err = json.Unmarshal(result, &match)
	if err != nil {
		return nil, err
	}
	return &match, nil
}

// GetMatchTimelineById is a function that returns a match timeline by match ID.
func (cli *RiotClient) GetMatchTimelineById(matchId string) (*response.MatchTimeline, error) {
	result, err := cli.riotClient.Get("/lol/match/v5/matches/:matchId/timeline", []request.ParamURL{
		{
			Key:   "matchId",
			Value: matchId,
		},
	})
	if err != nil {
		return nil, err
	}
	var matchTimeline response.MatchTimeline
	err = json.Unmarshal(result, &matchTimeline)
	if err != nil {
		return nil, err
	}
	return &matchTimeline, nil
}
