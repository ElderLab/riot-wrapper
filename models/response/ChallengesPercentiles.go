package response

import "encoding/json"

// ChallengesPercentiles is a map that represents the challenges percentiles.
type ChallengesPercentiles map[string]float64

// String is a function that returns the challenges percentiles in json format.
func (cp ChallengesPercentiles) String() string {
	b, err := json.MarshalIndent(cp, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// AllChallengesPercentiles is a map that represents the all challenges percentiles.
type AllChallengesPercentiles map[string]ChallengesPercentiles

// String is a function that returns the all challenges percentiles in json format.
func (acp AllChallengesPercentiles) String() string {
	b, err := json.MarshalIndent(acp, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
