package response

import "encoding/json"

// ChallengesLeaderboards is a struct that represents the challenges leaderboards
type ChallengesLeaderboards struct {
	Position int     `json:"position"`
	Puuid    string  `json:"puuid"`
	Value    float64 `json:"value"`
}

// String is a function that returns the challenges leaderboards in json format
func (cl ChallengesLeaderboards) String() string {
	b, err := json.MarshalIndent(cl, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
