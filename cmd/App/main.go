package main

import (
	"fmt"
	riotwrapper "github.com/ElderLab/riot-wrapper"
	"github.com/ElderLab/riot-wrapper/models/opts"
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
	//r1, errs := cliRiot.GetChampionRotation()
	//if len(errs) > 0 {
	//	panic(errs)
	//}
	//fmt.Println(r1)
	r2, err := cliRiot.GetAccountByRiotId("Sambre", "3005")
	if err != nil {
		panic(err)
	}
	fmt.Println(r2)
	r3, err := cliRiot.GetSummonerByPUUID(r2.Puuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(r3)

	r4, err := cliRiot.GetSummonerByAccountId(r3.AccountId)
	if err != nil {
		panic(err)
	}
	fmt.Println(r4)

	r5, err := cliRiot.GetMatchesIds(r2.Puuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(r5)
	r6, err := cliRiot.GetMatchById((*r5)[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(r6)
	r7, err := cliRiot.GetMatchTimelineById((*r5)[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(r7)
	r8, err := cliRiot.GetMatchesIdsWithOpts(r2.Puuid, opts.MatchesIdsOpts{
		Count: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(r8)

	r9, err := cliRiot.GetChallengerByQueue(riotwrapper.RankedSolo5x5)
	if err != nil {
		panic(err)
	}
	fmt.Println(r9)

	r10, err := cliRiot.GetAllChampionMastery(r2.Puuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(r10)
}
