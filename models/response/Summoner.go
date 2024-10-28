package response

import "encoding/json"

// Summoner is a struct that represents the summoner riot.
type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// String is a function that returns the summoner in json format.
func (s Summoner) String() string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
