package riot_wrapper

import "ElderLab/riot-wrapper/internal/request"

// GetMatchesIds is a function that returns a list of match IDs by PUUID.
func (cli *RiotClient) GetMatchesIds(puuid string) (string, []error) {
	result, err := cli.riotClient.Get("/lol/match/v5/matches/by-puuid/:puuid/ids", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return "", err
	}
	return string(result), nil
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
