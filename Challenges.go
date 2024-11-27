package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
	"strconv"
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
