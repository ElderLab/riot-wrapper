package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/opts"
	"github.com/ElderLab/riot-wrapper/models/response"
	"strconv"
)

type rank string

const (
	// MASTER is a rank type
	MASTER rank = "MASTER"
	// GRANDMASTER is a rank type
	GRANDMASTER rank = "GRANDMASTER"
	// CHALLENGER is a rank type
	CHALLENGER rank = "CHALLENGER"
)

// GetChallengesConfig returns the challenges configuration
func (cli *RiotClient) GetChallengesConfig() (*[]response.Challenges, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/config", []request.ParamURL{})
	if err != nil {
		return nil, err
	}
	var challenges []response.Challenges
	err = json.Unmarshal(result, &challenges)
	if err != nil {
		return nil, err
	}
	return &challenges, nil
}

// GetChallengesConfigById returns the challenges configuration by ID
func (cli *RiotClient) GetChallengesConfigById(id int) (*response.Challenges, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/:id/config", []request.ParamURL{
		{
			Key:   "id",
			Value: strconv.Itoa(id),
		},
	})
	if err != nil {
		return nil, err
	}
	var challenges response.Challenges
	err = json.Unmarshal(result, &challenges)
	if err != nil {
		return nil, err
	}
	return &challenges, nil
}

// GetChallengesPercentiles returns all challenges percentiles
func (cli *RiotClient) GetChallengesPercentiles() (*response.AllChallengesPercentiles, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/percentiles", []request.ParamURL{})
	if err != nil {
		return nil, err
	}
	var allChallengesPercentiles response.AllChallengesPercentiles
	err = json.Unmarshal(result, &allChallengesPercentiles)
	if err != nil {
		return nil, err
	}
	return &allChallengesPercentiles, nil
}

// GetChallengesPercentilesById returns the challenges percentiles by challenge ID
func (cli *RiotClient) GetChallengesPercentilesById(challengeId int) (*response.ChallengesPercentiles, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/:challengeId/percentiles", []request.ParamURL{
		{
			Key:   "challengeId",
			Value: strconv.Itoa(challengeId),
		},
	})
	if err != nil {
		return nil, err
	}
	var challengesPercentiles response.ChallengesPercentiles
	err = json.Unmarshal(result, &challengesPercentiles)
	if err != nil {
		return nil, err
	}
	return &challengesPercentiles, nil
}

// GetChallengesLeaderboards returns the challenges leaderboards
func (cli *RiotClient) GetChallengesLeaderboards(challengeId int, level rank) (*[]response.ChallengesLeaderboards, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/:challengeId/leaderboards/by-level/:level", []request.ParamURL{
		{
			Key:   "challengeId",
			Value: strconv.Itoa(challengeId),
		},
		{
			Key:   "level",
			Value: string(level),
		},
	})
	if err != nil {
		return nil, err
	}
	var challengesLeaderboards []response.ChallengesLeaderboards
	err = json.Unmarshal(result, &challengesLeaderboards)
	if err != nil {
		return nil, err
	}
	return &challengesLeaderboards, nil
}

// GetChallengesLeaderboardsWithOpts returns the challenges leaderboards with options
func (cli *RiotClient) GetChallengesLeaderboardsWithOpts(challengeId int, level rank, opts opts.ChallengesLeaderboardsOpts) (*[]response.ChallengesLeaderboards, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/challenges/:challengeId/leaderboards/by-level/:level", []request.ParamURL{
		{
			Key:   "challengeId",
			Value: strconv.Itoa(challengeId),
		},
		{
			Key:   "level",
			Value: string(level),
		},
		{
			Key:   "limit",
			Value: strconv.Itoa(opts.Limit),
		},
	})
	if err != nil {
		return nil, err
	}
	var challengesLeaderboards []response.ChallengesLeaderboards
	err = json.Unmarshal(result, &challengesLeaderboards)
	if err != nil {
		return nil, err
	}
	return &challengesLeaderboards, nil
}

func (cli *RiotClient) GetChallengesPlayerProgression(puuid string) (*response.ChallengesProgression, error) {
	result, err := cli.lolClient.Get("/lol/challenges/v1/player-data/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	if err != nil {
		return nil, err
	}
	var challengesProgression response.ChallengesProgression
	err = json.Unmarshal(result, &challengesProgression)
	if err != nil {
		return nil, err
	}
	return &challengesProgression, nil
}
