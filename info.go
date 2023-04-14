package webflow

import "time"

type Info struct {
	Id          string       `json:"_id,omitempty"`
	GrantType   string       `json:"grantType,omitempty"`
	Status      string       `json:"status,omitempty"`
	RateLimit   int          `json:"rateLimit,omitempty"`
	Sites       []string     `json:"sites,omitempty"`
	Orgs        []string     `json:"orgs,omitempty"`
	Workspaces  []string     `json:"workspaces,omitempty"`
	Users       []string     `json:"users,omitempty"`
	CreatedOn   time.Time    `json:"createdOn,omitempty"`
	LastUsed    time.Time    `json:"lastUsed,omitempty"`
	Application *Application `json:"application,omitempty"`
}

func (c *Client) AuthorizedInfo() (*Info, error) {
	info := new(Info)
	if err := c.Get("info/", nil, nil, info); err != nil {
		return nil, err
	}
	return info, nil
}
