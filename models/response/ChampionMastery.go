package response

import "encoding/json"

// ChampionMastery is a struct that represents the champion mastery of a summoner.
type ChampionMastery struct {
	Puuid                        string `json:"puuid"`
	ChampionId                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	MarkRequiredForNextLevel     int    `json:"markRequiredForNextLevel"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionSeasonMilestone      int    `json:"championSeasonMilestone"`
	NextSeasonMilestone          struct {
		RequireGradeCounts struct {
			S int `json:"S-,omitempty"`
			A int `json:"A-,omitempty"`
			B int `json:"B-,omitempty"`
			C int `json:"C-,omitempty"`
		} `json:"requireGradeCounts"`
		RewardMarks  int  `json:"rewardMarks"`
		Bonus        bool `json:"bonus"`
		RewardConfig struct {
			RewardValue   string `json:"rewardValue"`
			RewardType    string `json:"rewardType"`
			MaximumReward int    `json:"maximumReward"`
		} `json:"rewardConfig,omitempty"`
		TotalGamesRequires int `json:"totalGamesRequires"`
	} `json:"nextSeasonMilestone"`
	MilestoneGrades []string `json:"milestoneGrades,omitempty"`
}

// String is a function that returns the champion mastery in json format.
func (cm ChampionMastery) String() string {
	b, err := json.MarshalIndent(cm, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
