package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
)

// GetChallengerByQueue is a function that returns the challenger league for a given queue.
func (cli *RiotClient) GetChallengerByQueue(queue queue) (*response.ChallengerLeague, error) {
	result, err := cli.lolClient.Get("/lol/league/v4/challengerleagues/by-queue/"+string(queue), []request.ParamURL{})
	if err != nil {
		return nil, err
	}

	var challengerLeague response.ChallengerLeague
	err = json.Unmarshal(result, &challengerLeague)
	if err != nil {
		return nil, err
	}
	return &challengerLeague, nil
}
