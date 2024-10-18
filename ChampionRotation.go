package riot_wrapper

import "ElderLab/riot-wrapper/internal/request"

// GetChampionRotation is a function that returns the free champion rotation for the current week.
func (cli *RiotClient) GetChampionRotation() (string, []error) {

	result, err := cli.lolClient.Get("/lol/platform/v3/champion-rotations", []request.ParamURL{})

	if err != nil {
		return "", err
	}
	return string(result), nil
}
