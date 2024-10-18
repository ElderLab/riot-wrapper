package main

import (
	riot_wrapper "ElderLab/riot-wrapper"
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
	cliRiot := riot_wrapper.NewRiotClient(apiKey)
	r1, err := cliRiot.GetChampionRotation()
	if len(err) > 0 {
		panic(err)
	}
	fmt.Println(r1)
	r2, err := cliRiot.GetAccountByRiotId("Sambre", "3005")
	if len(err) > 0 {
		panic(err)
	}
	fmt.Println(r2)
}
