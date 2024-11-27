package response

import "encoding/json"

type ChallengesProgression struct {
	TotalPoints struct {
		Level      string  `json:"level"`
		Current    int     `json:"current"`
		Max        int     `json:"max"`
		Percentile float64 `json:"percentile"`
	} `json:"totalPoints"`
	CategoryPoints struct {
		TEAMWORK struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"TEAMWORK"`
		IMAGINATION struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"IMAGINATION"`
		COLLECTION struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"COLLECTION"`
		VETERANCY struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"VETERANCY"`
		EXPERTISE struct {
			Level      string  `json:"level"`
			Current    int     `json:"current"`
			Max        int     `json:"max"`
			Percentile float64 `json:"percentile"`
		} `json:"EXPERTISE"`
	} `json:"categoryPoints"`
	Challenges []struct {
		ChallengeId    int     `json:"challengeId"`
		Percentile     float64 `json:"percentile"`
		Level          string  `json:"level"`
		Value          float64 `json:"value"`
		AchievedTime   int64   `json:"achievedTime,omitempty"`
		Position       int     `json:"position,omitempty"`
		PlayersInLevel int     `json:"playersInLevel,omitempty"`
	} `json:"challenges"`
	Preferences struct {
		BannerAccent             string `json:"bannerAccent"`
		Title                    string `json:"title"`
		ChallengeIds             []int  `json:"challengeIds"`
		CrestBorder              string `json:"crestBorder"`
		PrestigeCrestBorderLevel int    `json:"prestigeCrestBorderLevel"`
	} `json:"preferences"`
}

// String is a function that returns the ChallengesProgression in a json format
func (cp ChallengesProgression) String() string {
	b, err := json.MarshalIndent(cp, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
