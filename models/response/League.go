package response

import "encoding/json"

// League is a struct that represents a tier league.
type League struct {
	Tier     string `json:"tier"`
	LeagueId string `json:"leagueId"`
	Queue    string `json:"queue"`
	Name     string `json:"name"`
	Entries  []struct {
		SummonerId   string `json:"summonerId"`
		LeaguePoints int    `json:"leaguePoints"`
		Rank         string `json:"rank"`
		Wins         int    `json:"wins"`
		Losses       int    `json:"losses"`
		Veteran      bool   `json:"veteran"`
		Inactive     bool   `json:"inactive"`
		FreshBlood   bool   `json:"freshBlood"`
		HotStreak    bool   `json:"hotStreak"`
	} `json:"entries"`
}

// String is a function that returns  the league in json format.
func (c League) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
