package response

import "encoding/json"

type PlayerRank struct {
	LeagueId     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerId   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

// String is a function that returns the player rank in json format.
func (pr PlayerRank) String() string {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
