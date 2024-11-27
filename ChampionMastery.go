package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/opts"
	"github.com/ElderLab/riot-wrapper/models/response"
	"strconv"
)

// GetAllChampionMastery is a function that returns all the champion mastery of a summoner by PUUID.
func (cli *RiotClient) GetAllChampionMastery(puuid string) (*[]response.ChampionMastery, error) {
	result, err := cli.lolClient.Get("/lol/champion-mastery/v4/champion-masteries/by-puuid/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return nil, err
	}
	var championMastery []response.ChampionMastery
	err = json.Unmarshal(result, &championMastery)
	if err != nil {
		return nil, err
	}
	return &championMastery, nil
}

// GetChampionMasteryById is a function that returns the champion mastery of a summoner by PUUID and champion ID.
func (cli *RiotClient) GetChampionMasteryById(puuid string, championId int) (*response.ChampionMastery, error) {
	result, err := cli.lolClient.Get("/lol/champion-mastery/v4/champion-masteries/by-puuid/:puuid/by-champion/:championId", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
		{
			Key:   "championId",
			Value: strconv.Itoa(championId),
		},
	})
	if err != nil {
		return nil, err
	}
	var championMastery response.ChampionMastery
	err = json.Unmarshal(result, &championMastery)
	if err != nil {
		return nil, err
	}
	return &championMastery, nil
}

// GetChampionMastery is a function that returns the champion mastery of a summoner by PUUID.
func (cli *RiotClient) GetChampionMastery(puuid string) (*[]response.ChampionMastery, error) {
	result, err := cli.lolClient.Get("/lol/champion-mastery/v4/champion-masteries/by-puuid/:puuid/top", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return nil, err
	}
	var championMastery []response.ChampionMastery
	err = json.Unmarshal(result, &championMastery)
	if err != nil {
		return nil, err
	}
	return &championMastery, nil
}

// GetChampionMasteryWithOpts is a function that returns the champion mastery of a summoner by PUUID with options.
func (cli *RiotClient) GetChampionMasteryWithOpts(puuid string, opts opts.ChampionMasteryOpts) (*[]response.ChampionMastery, error) {
	result, err := cli.lolClient.Get("/lol/champion-mastery/v4/champion-masteries/by-puuid/:puuid/top", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
		{
			Key:   "count",
			Value: strconv.Itoa(opts.Count),
		},
	})
	if err != nil {
		return nil, err
	}
	var championMastery []response.ChampionMastery
	err = json.Unmarshal(result, &championMastery)
	if err != nil {
		return nil, err
	}
	return &championMastery, nil
}

// GetChampionMasteryScore is a function that returns the champion mastery score of a summoner by PUUID.
func (cli *RiotClient) GetChampionMasteryScore(puuid string) (int, error) {
	result, err := cli.lolClient.Get("/lol/champion-mastery/v4/scores/by-puuid/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return 0, err
	}
	score, err := strconv.Atoi(string(result))
	if err != nil {
		return 0, err
	}
	return score, nil
}
