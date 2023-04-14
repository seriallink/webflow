package webflow

type Client struct {
	token   string
	version string
}

func NewClient(token string) *Client {
	return &Client{
		token:   token,
		version: apiVersion,
	}
}

func (c *Client) GetToken() string {
	return c.token
}

func (c *Client) GetVersion() string {
	return c.version
}
