package request

// Client is a struct that contains the base URL and the X-Riot-Token.
type Client struct {
	BaseURL    string
	XRiotToken string
}

// NewClient is a function that returns a new Client.
func NewClient(BaseURL, XRiotToken string) *Client {
	return &Client{
		BaseURL:    BaseURL,
		XRiotToken: XRiotToken,
	}
}
