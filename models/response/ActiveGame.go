package response

import "encoding/json"

// ActiveGame is a struct that represents the current game information.
type ActiveGame struct {
	GameId            int64  `json:"gameId"`
	MapId             int    `json:"mapId"`
	GameMode          string `json:"gameMode"`
	GameType          string `json:"gameType"`
	GameQueueConfigId int    `json:"gameQueueConfigId"`
	Participants      []struct {
		Puuid                    string        `json:"puuid"`
		TeamId                   int           `json:"teamId"`
		Spell1Id                 int           `json:"spell1Id"`
		Spell2Id                 int           `json:"spell2Id"`
		ChampionId               int           `json:"championId"`
		ProfileIconId            int           `json:"profileIconId"`
		RiotId                   string        `json:"riotId"`
		Bot                      bool          `json:"bot"`
		SummonerId               string        `json:"summonerId"`
		GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
		Perks                    struct {
			PerkIds      []int `json:"perkIds"`
			PerkStyle    int   `json:"perkStyle"`
			PerkSubStyle int   `json:"perkSubStyle"`
		} `json:"perks"`
	} `json:"participants"`
	Observers struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	PlatformId      string `json:"platformId"`
	BannedChampions []struct {
		ChampionId int `json:"championId"`
		TeamId     int `json:"teamId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
	GameStartTime int64 `json:"gameStartTime"`
	GameLength    int   `json:"gameLength"`
}

// String is a function that returns the active game in json format.
func (ag ActiveGame) String() string {
	b, err := json.MarshalIndent(ag, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
