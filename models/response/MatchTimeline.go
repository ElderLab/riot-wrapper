package response

import "encoding/json"

// MatchTimeline is a struct that represents the timeline of a match.
type MatchTimeline struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchId      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult string `json:"endOfGameResult"`
		FrameInterval   int    `json:"frameInterval"`
		Frames          []struct {
			Events []struct {
				RealTimestamp           int64  `json:"realTimestamp,omitempty"`
				Timestamp               int    `json:"timestamp"`
				Type                    string `json:"type"`
				LevelUpType             string `json:"levelUpType,omitempty"`
				ParticipantId           int    `json:"participantId,omitempty"`
				SkillSlot               int    `json:"skillSlot,omitempty"`
				ItemId                  int    `json:"itemId,omitempty"`
				CreatorId               int    `json:"creatorId,omitempty"`
				WardType                string `json:"wardType,omitempty"`
				Level                   int    `json:"level,omitempty"`
				AssistingParticipantIds []int  `json:"assistingParticipantIds,omitempty"`
				Bounty                  int    `json:"bounty,omitempty"`
				KillStreakLength        int    `json:"killStreakLength,omitempty"`
				KillerId                int    `json:"killerId,omitempty"`
				Position                struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"position,omitempty"`
				ShutdownBounty    int `json:"shutdownBounty,omitempty"`
				VictimDamageDealt []struct {
					Basic          bool   `json:"basic"`
					MagicDamage    int    `json:"magicDamage"`
					Name           string `json:"name"`
					ParticipantId  int    `json:"participantId"`
					PhysicalDamage int    `json:"physicalDamage"`
					SpellName      string `json:"spellName"`
					SpellSlot      int    `json:"spellSlot"`
					TrueDamage     int    `json:"trueDamage"`
					Type           string `json:"type"`
				} `json:"victimDamageDealt,omitempty"`
				VictimDamageReceived []struct {
					Basic          bool   `json:"basic"`
					MagicDamage    int    `json:"magicDamage"`
					Name           string `json:"name"`
					ParticipantId  int    `json:"participantId"`
					PhysicalDamage int    `json:"physicalDamage"`
					SpellName      string `json:"spellName"`
					SpellSlot      int    `json:"spellSlot"`
					TrueDamage     int    `json:"trueDamage"`
					Type           string `json:"type"`
				} `json:"victimDamageReceived,omitempty"`
				VictimId        int    `json:"victimId,omitempty"`
				KillType        string `json:"killType,omitempty"`
				MultiKillLength int    `json:"multiKillLength,omitempty"`
				LaneType        string `json:"laneType,omitempty"`
				TeamId          int    `json:"teamId,omitempty"`
				AfterId         int    `json:"afterId,omitempty"`
				BeforeId        int    `json:"beforeId,omitempty"`
				GoldGain        int    `json:"goldGain,omitempty"`
				KillerTeamId    int    `json:"killerTeamId,omitempty"`
				MonsterSubType  string `json:"monsterSubType,omitempty"`
				MonsterType     string `json:"monsterType,omitempty"`
				BuildingType    string `json:"buildingType,omitempty"`
				TowerType       string `json:"towerType,omitempty"`
				ActualStartTime int    `json:"actualStartTime,omitempty"`
				GameId          int64  `json:"gameId,omitempty"`
				WinningTeam     int    `json:"winningTeam,omitempty"`
			} `json:"events"`
			ParticipantFrames map[string]ParticipantFrame `json:"participantFrames"`
			Timestamp         int                         `json:"timestamp"`
		} `json:"frames"`
		GameId       int64 `json:"gameId"`
		Participants []struct {
			ParticipantId int    `json:"participantId"`
			Puuid         string `json:"puuid"`
		} `json:"participants"`
	} `json:"info"`
}

type ParticipantFrame struct {
	ChampionStats struct {
		AbilityHaste         int `json:"abilityHaste"`
		AbilityPower         int `json:"abilityPower"`
		Armor                int `json:"armor"`
		ArmorPen             int `json:"armorPen"`
		ArmorPenPercent      int `json:"armorPenPercent"`
		AttackDamage         int `json:"attackDamage"`
		AttackSpeed          int `json:"attackSpeed"`
		BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
		BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
		CcReduction          int `json:"ccReduction"`
		CooldownReduction    int `json:"cooldownReduction"`
		Health               int `json:"health"`
		HealthMax            int `json:"healthMax"`
		HealthRegen          int `json:"healthRegen"`
		Lifesteal            int `json:"lifesteal"`
		MagicPen             int `json:"magicPen"`
		MagicPenPercent      int `json:"magicPenPercent"`
		MagicResist          int `json:"magicResist"`
		MovementSpeed        int `json:"movementSpeed"`
		Omnivamp             int `json:"omnivamp"`
		PhysicalVamp         int `json:"physicalVamp"`
		Power                int `json:"power"`
		PowerMax             int `json:"powerMax"`
		PowerRegen           int `json:"powerRegen"`
		SpellVamp            int `json:"spellVamp"`
	} `json:"championStats"`
	CurrentGold int `json:"currentGold"`
	DamageStats struct {
		MagicDamageDone               int `json:"magicDamageDone"`
		MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
		MagicDamageTaken              int `json:"magicDamageTaken"`
		PhysicalDamageDone            int `json:"physicalDamageDone"`
		PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
		PhysicalDamageTaken           int `json:"physicalDamageTaken"`
		TotalDamageDone               int `json:"totalDamageDone"`
		TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
		TotalDamageTaken              int `json:"totalDamageTaken"`
		TrueDamageDone                int `json:"trueDamageDone"`
		TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
		TrueDamageTaken               int `json:"trueDamageTaken"`
	} `json:"damageStats"`
	GoldPerSecond       int `json:"goldPerSecond"`
	JungleMinionsKilled int `json:"jungleMinionsKilled"`
	Level               int `json:"level"`
	MinionsKilled       int `json:"minionsKilled"`
	ParticipantId       int `json:"participantId"`
	Position            struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	TimeEnemySpentControlled int `json:"timeEnemySpentControlled"`
	TotalGold                int `json:"totalGold"`
	Xp                       int `json:"xp"`
}

// String is a function that returns the match timeline in json format.
func (m MatchTimeline) String() string {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
