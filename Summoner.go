package riot_wrapper

import "ElderLab/riot-wrapper/internal/request"

func (cli *RiotClient) GetSummonerByPUUID(puuid string) (string, []error) {
	result, err := cli.lolClient.Get("/lol/summoner/v4/summoners/by-puuid/:puuid", []request.ParamURL{
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

func (cli *RiotClient) GetSummonerByAccountId(accountId string) (string, []error) {
	result, err := cli.lolClient.Get("/lol/summoner/v4/summoners/by-account/:accountId", []request.ParamURL{
		{
			Key:   "accountId",
			Value: accountId,
		},
	})
	if err != nil {
		return "", err
	}
	return string(result), nil
}
