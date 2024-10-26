package request

// Client is a struct that contains the base URL and the X-Riot-Token.
type Client struct {
	baseURL    string
	xRiotToken string
	lastStatus int
}

// NewClient is a function that returns a new Client.
func NewClient(BaseURL, XRiotToken string) *Client {
	return &Client{
		baseURL:    BaseURL,
		xRiotToken: XRiotToken,
	}
}

func (c *Client) GetLastStatus() int {
	return c.lastStatus
}
