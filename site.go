package webflow

import "time"

type Site struct {
	Id            string    `json:"_id,omitempty"`
	Name          string    `json:"name,omitempty"`
	ShortName     string    `json:"shortName,omitempty"`
	PreviewUrl    string    `json:"previewUrl,omitempty"`
	Timezone      string    `json:"timezone,omitempty"`
	Database      string    `json:"database,omitempty"`
	CreatedOn     time.Time `json:"createdOn,omitempty"`
	LastPublished time.Time `json:"lastPublished,omitempty"`
}

func (c *Client) ListSites() (sites []Site, err error) {
	err = c.Get("sites/", nil, nil, &sites)
	return
}

func (c *Client) GetSite(siteId string) (*Site, error) {
	site := new(Site)
	if err := c.Get("sites/"+siteId, nil, nil, site); err != nil {
		return nil, err
	}
	return site, nil
}

func (c *Client) PublishSite(siteId string, domains []string) (*QueueResponse, error) {
	queue := new(QueueResponse)
	params := Params{
		"domains": domains,
	}
	if err := c.Post("sites/"+siteId+"/publish", params, nil, queue); err != nil {
		return nil, err
	}
	return queue, nil
}
