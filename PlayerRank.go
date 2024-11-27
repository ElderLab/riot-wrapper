package riot_wrapper

import (
	"encoding/json"
	"github.com/ElderLab/riot-wrapper/internal/request"
	"github.com/ElderLab/riot-wrapper/models/response"
)

func (cli *RiotClient) GetPlayerRankBySummonerId(summonerId string) ([]response.PlayerRank, error) {
	result, err := cli.lolClient.Get("/lol/league/v4/entries/by-summoner/:summonerId", []request.ParamURL{
		{
			Key:   "summonerId",
			Value: summonerId,
		},
	})
	if err != nil {
		return nil, err
	}

	var playerRank []response.PlayerRank
	err = json.Unmarshal(result, &playerRank)
	if err != nil {
		return nil, err
	}
	return playerRank, nil
}
