package response

import "encoding/json"

// Challenges is a struct that represents the challenges configuration.
type Challenges struct {
	Id             int `json:"id"`
	LocalizedNames struct {
		ArAE struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"ar_AE"`
		CsCZ struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"cs_CZ"`
		DeDE struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"de_DE"`
		ElGR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"el_GR"`
		EnAU struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"en_AU"`
		EnGB struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"en_GB"`
		EnPH struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"en_PH"`
		EnSG struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"en_SG"`
		EnUS struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"en_US"`
		EsAR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"es_AR"`
		EsES struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"es_ES"`
		EsMX struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"es_MX"`
		FrFR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"fr_FR"`
		HuHU struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"hu_HU"`
		ItIT struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"it_IT"`
		JaJP struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"ja_JP"`
		KoKR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"ko_KR"`
		PlPL struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"pl_PL"`
		PtBR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"pt_BR"`
		RoRO struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"ro_RO"`
		RuRU struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"ru_RU"`
		ThTH struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"th_TH"`
		TrTR struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"tr_TR"`
		ViVN struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"vi_VN"`
		ZhCN struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"zh_CN"`
		ZhMY struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"zh_MY"`
		ZhTW struct {
			Description      string `json:"description"`
			Name             string `json:"name"`
			ShortDescription string `json:"shortDescription"`
		} `json:"zh_TW"`
	} `json:"localizedNames"`
	State       string `json:"state"`
	Leaderboard bool   `json:"leaderboard"`
	Thresholds  struct {
		BRONZE      float64 `json:"BRONZE,omitempty"`
		PLATINUM    float64 `json:"PLATINUM"`
		DIAMOND     float64 `json:"DIAMOND"`
		GRANDMASTER float64 `json:"GRANDMASTER,omitempty"`
		MASTER      float64 `json:"MASTER"`
		GOLD        float64 `json:"GOLD"`
		SILVER      float64 `json:"SILVER,omitempty"`
		IRON        float64 `json:"IRON,omitempty"`
		CHALLENGER  float64 `json:"CHALLENGER,omitempty"`
	} `json:"thresholds"`
}

// String is a function that returns the challenges configuration in json format.
func (c Challenges) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
