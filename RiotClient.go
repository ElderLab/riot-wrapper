package riot_wrapper

import "ElderLab/riot-wrapper/internal/request"

// RiotClient is a struct that contains the RiotClient.
type RiotClient struct {
	riotClient *request.Client
	lolClient  *request.Client
}

// NewRiotClient is a function that returns a new RiotClient.
func NewRiotClient(Token string) *RiotClient {
	return &RiotClient{
		riotClient: request.NewClient("https://europe.api.riotgames.com", Token),
		lolClient:  request.NewClient("https://euw1.api.riotgames.com", Token),
	}
}
