package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
)

func (cli *RiotClient) GetActiveGame(puuid string) (*response.ActiveGame, error) {
	result, err := cli.lolClient.Get("/lol/spectator/v5/active-games/by-summoner/:puuid", []request.ParamURL{
		{
			Key:   "puuid",
			Value: puuid,
		},
	})
	// If the player is not in a game, the API will return a 404 error
	if err != nil && err.Error() == "Status : 404" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var activeGame response.ActiveGame
	err = json.Unmarshal(result, &activeGame)
	if err != nil {
		return nil, err
	}
	return &activeGame, nil
}
