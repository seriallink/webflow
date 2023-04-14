package webflow

type Domain struct {
	Id   string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) ListDomains(siteId string) (domains []Domain, err error) {
	err = c.Get("sites/"+siteId+"/domains/", nil, nil, &domains)
	return
}
