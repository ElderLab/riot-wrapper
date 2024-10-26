package main

import (
	riotwrapper "ElderLab/riot-wrapper"
	"encoding/json"
	"fmt"
	"github.com/tot0p/env"
)

func init() {
	err := env.Load()
	if err != nil {
		return
	}
}

func main() {
	apiKey := env.Get("RIOT_API_KEY")
	cliRiot := riotwrapper.NewRiotClient(apiKey)
	r1, errs := cliRiot.GetChampionRotation()
	if len(errs) > 0 {
		panic(errs)
	}
	fmt.Println(r1)
	r2, errs := cliRiot.GetAccountByRiotId("Sambre", "3005")
	if len(errs) > 0 {
		panic(errs)
	}
	fmt.Println(r2)
	r3, errs := cliRiot.GetSummonerByPUUID(r2.Puuid)
	if len(errs) > 0 {
		panic(errs)
	}
	fmt.Println(r3)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(r3), &result)
	if err != nil {
		panic(err)
	}
	r4, errs := cliRiot.GetSummonerByAccountId(result["accountId"].(string))
	if len(errs) > 0 {
		panic(err)
	}
	fmt.Println(r4)

	r5, errs := cliRiot.GetMatchesIds(r2.Puuid)
	if len(errs) > 0 {
		panic(err)
	}

	fmt.Println(r5)
}
